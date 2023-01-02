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
			log.Ctx(ctx).Panic().Err(err).Interface("v.Index", v.Index).Msg("InsertValidatorsInMemDb")
			panic(err)
		}
	}
	txn.Commit()
}

func ReadOnlyValidatorFromInMemDb(ctx context.Context, v Validator) Validator {
	txn := ValidatorInMemDB.Txn(false)
	defer txn.Abort()
	raw, err := txn.First(ValidatorsTable, "validator_index", v.Index)
	if err != nil {
		log.Ctx(ctx).Panic().Err(err).Interface("v.Index", v.Index).Msg("ReadOnlyValidatorFromInMemDb")
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
