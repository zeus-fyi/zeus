package aegis_inmemdbs

import (
	"context"

	"github.com/hashicorp/go-memdb"
	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/gochain/web3/accounts"
	"github.com/zeus-fyi/zeus/pkg/crypto/ecdsa"
)

var EcdsaAccountsInMemDB *memdb.MemDB

const EcdsaAccountsTable = "ecdsa_accounts"

// hex formatted, eg 0x prefixed
type ecdsaAccount struct {
	publicKey  string
	privateKey string
}

func InsertEcdsaAccounts(ctx context.Context, accounts []ecdsa.Account) {
	txn := EcdsaAccountsInMemDB.Txn(true)
	for _, acc := range accounts {
		insertAccount := ecdsaAccount{acc.PublicKey(), acc.PrivateKey()}
		if err := txn.Insert("ecdsa_accounts", insertAccount); err != nil {
			log.Ctx(ctx).Panic().Err(err).Interface("public_key", acc.PublicKey()).Msg("InsertEcdsaAccounts")
			panic(err)
		}
	}
	txn.Commit()
}

func ReadOnlyEcdsaAccountFromInMemDb(ctx context.Context, a ecdsa.Account) ecdsa.Account {
	txn := EcdsaAccountsInMemDB.Txn(false)
	defer txn.Abort()
	raw, err := txn.First(EcdsaAccountsTable, "public_key", a.PublicKey())
	if err != nil {
		log.Ctx(ctx).Panic().Err(err).Interface("public_key", a.PublicKey()).Msg("ReadOnlyEcdsaAccountFromInMemDb")
		panic(err)
	}
	txn.Commit()
	return convertInMemEcdsaAccountToAccount(ctx, raw.(ecdsaAccount))
}

func convertInMemEcdsaAccountToAccount(ctx context.Context, e ecdsaAccount) ecdsa.Account {
	acc, err := accounts.ParsePrivateKey(e.privateKey)
	if err != nil {
		log.Ctx(ctx).Panic().Err(err).Interface("public_key", e.publicKey).Msg("convertInMemEcdsaAccountToAccount")
		panic(err)
	}
	return ecdsa.Account{Account: acc}
}

func InitEcdsaAccountsDB() {
	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			EcdsaAccountsTable: {
				Name: EcdsaAccountsTable,
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "publicKey"},
					},
					"public_key": {
						Name:    "public_key",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "publicKey"},
					},
					"private_key": {
						Name:    "private_key",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "privateKey"},
					},
				},
			},
		},
	}
	// Create a new database
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		log.Panic().Err(err).Interface("table", EcdsaAccountsTable).Msg("InitEcdsaAccountsDB")
		panic(err)
	}
	EcdsaAccountsInMemDB = db
}
