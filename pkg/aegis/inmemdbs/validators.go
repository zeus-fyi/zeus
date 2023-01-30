package aegis_inmemdbs

import (
	"context"

	"github.com/hashicorp/go-memdb"
	"github.com/rs/zerolog/log"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
)

var ValidatorInMemDB *memdb.MemDB

type Validator struct {
	bls_signer.EthBLSAccount
}

const ValidatorsTable = "validators"

type inMemValidator struct {
	Index     string
	PublicKey []byte
	SecretKey []byte
}

func NewValidator(blsKey bls_signer.EthBLSAccount) Validator {
	v := Validator{
		EthBLSAccount: blsKey,
	}
	return v
}

func InsertValidatorsInMemDb(ctx context.Context, vs []Validator) {
	txn := ValidatorInMemDB.Txn(true)
	for _, v := range vs {
		insertV := inMemValidator{
			Index:     v.PublicKeyString(),
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
						Indexer: &memdb.StringFieldIndex{Field: "Index"},
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
