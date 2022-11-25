package aegis

import (
	"github.com/hashicorp/go-memdb"
	bls_signer "github.com/zeus-fyi/zeus/pkg/crypto/bls"
)

var ValidatorInMemDB *memdb.MemDB

type Validator struct {
	Index string
	bls_signer.KeyBLS
}

func InsertValidators(vs []Validator) {
	txn := ValidatorInMemDB.Txn(true)
	for _, v := range vs {
		if err := txn.Insert("validator", v); err != nil {
			panic(err)
		}
	}
	txn.Commit()
}

func InitValidatorDB() {
	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"validator": &memdb.TableSchema{
				Name: "validator",
				Indexes: map[string]*memdb.IndexSchema{
					"index": &memdb.IndexSchema{
						Name:    "index",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Index"},
					},
					"pubKey": &memdb.IndexSchema{
						Name:    "pubKey",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "PubKey"},
					},
					"privKey": &memdb.IndexSchema{
						Name:    "privKey",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "PrivKey"},
					},
				},
			},
		},
	}

	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
	ValidatorInMemDB = db
}
