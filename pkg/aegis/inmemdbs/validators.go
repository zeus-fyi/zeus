package aegis_inmemdbs

import (
	"context"

	"github.com/hashicorp/go-memdb"
	"github.com/rs/zerolog/log"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
)

var ValidatorInMemDB *memdb.MemDB

type Validator struct {
	Index int
	bls_signer.EthBLSAccount
}

const ValidatorsTable = "validators"

type inMemValidator struct {
	Index     int
	PublicKey []byte
	SecretKey []byte
}

func NewValidator(index int, blsKey bls_signer.EthBLSAccount) Validator {
	v := Validator{
		Index:         index,
		EthBLSAccount: blsKey,
	}
	return v
}

func InsertValidatorsInMemDb(ctx context.Context, vs []Validator) {
	txn := ValidatorInMemDB.Txn(true)
	for _, v := range vs {
		insertV := inMemValidator{
			Index:     v.Index,
			PublicKey: v.PublicKey().Marshal(),
			SecretKey: v.BLSPrivateKey.Marshal(),
		}
		if err := txn.Insert("validators", insertV); err != nil {
			log.Ctx(ctx).Panic().Err(err).Interface("v.PublicKey", v.PublicKey()).Msg("InsertValidatorsInMemDb")
			panic(err)
		}
	}
	txn.Commit()
}

type EthereumBLSKeySignatureRequests struct {
	Map map[string]EthereumBLSKeySignatureRequest
}

type EthereumBLSKeySignatureRequest struct {
	Message string `json:"message"`
}

type EthereumBLSKeySignatureResponses struct {
	Map map[string]EthereumBLSKeySignatureResponse
}

type EthereumBLSKeySignatureResponse struct {
	Signature string `json:"signature"`
}

func SignValidatorMessagesFromInMemDb(ctx context.Context, signReqs EthereumBLSKeySignatureRequests) (EthereumBLSKeySignatureResponses, error) {
	resp := make(map[string]EthereumBLSKeySignatureResponse)
	batchResp := EthereumBLSKeySignatureResponses{
		Map: resp,
	}
	if len(signReqs.Map) == 0 {
		return batchResp, nil
	}
	txn := ValidatorInMemDB.Txn(false)
	defer txn.Abort()
	it, err := txn.Get("validators", "id")
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("SignValidatorMessagesFromInMemDb")
		return batchResp, err
	}

	tmp := make(map[string]Validator)
	for obj := it.Next(); obj != nil; obj = it.Next() {
		inMemDB := obj.(inMemValidator)
		v := bls_signer.NewEthSignerBLSFromExistingKeyBytes(inMemDB.SecretKey)
		pubkey := bls_signer.ConvertBytesToString(v.PublicKey().Marshal())
		tmp[pubkey] = NewValidator(inMemDB.Index, v)
	}
	txn.Commit()
	for _, v := range tmp {
		pubkey := v.PublicKeyString()
		msg, ok := signReqs.Map[pubkey]
		if ok {
			sig := v.Sign([]byte(msg.Message)).Marshal()
			batchResp.Map[pubkey] = EthereumBLSKeySignatureResponse{bls_signer.ConvertBytesToString(sig)}
		}
	}
	if len(batchResp.Map) != len(signReqs.Map) {
		log.Ctx(ctx).Warn().Msg("SignValidatorMessagesFromInMemDb, did not contain all expected validator signatures")
	}
	return batchResp, nil
}

func ReadOnlyValidatorFromInMemDb(ctx context.Context, pubkey string) Validator {
	txn := ValidatorInMemDB.Txn(false)
	defer txn.Abort()
	raw, err := txn.First(ValidatorsTable, "public_key", pubkey)
	if err != nil {
		log.Ctx(ctx).Panic().Err(err).Interface("v.public_key", pubkey).Msg("ReadOnlyValidatorFromInMemDb")
		panic(err)
	}
	txn.Commit()
	return raw.(Validator)
}

func InitValidatorDB() {
	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			ValidatorsTable: {
				Name: ValidatorsTable,
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "Index"},
					},
					"validator_index": {
						Name:    "validator_index",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "Index"},
					},
					"public_key": {
						Name:    "public_key",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "PublicKey"},
					},
					"secret_key": {
						Name:    "secret_key",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "SecretKey"},
					},
				},
			},
		},
	}
	// Create a new database
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		log.Panic().Err(err).Interface("table", ValidatorsTable).Msg("InitValidatorDB")
		panic(err)
	}
	ValidatorInMemDB = db
}
