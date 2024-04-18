package apps

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

type Db struct {
	Pgpool *pgxpool.Pool
	DbMap  map[string]*pgxpool.Pool
}

type Model interface {
	GetRowValues() RowValues
	GetManyRowValuesFlattened() RowValues
	GetManyRowValues() RowEntries
}

type RowValues []interface{}

type RowEntries struct {
	Rows []RowValues
}

var ConnStr string
var Pg Db

func (d *Db) InitPG(ctx context.Context, pgConnStr string) *pgxpool.Pool {
	Pg.DbMap = make(map[string]*pgxpool.Pool)
	Pg.Pgpool = d.InitAdditionalPG(ctx, "default", pgConnStr)
	return Pg.Pgpool
}

func (d *Db) InitAdditionalPG(ctx context.Context, name, pgConnStr string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(pgConnStr)
	if err != nil {
		log.Info().Msg("Zeus: InitPG failed to parse config to database")
		panic(err)
	}
	ConnStr = config.ConnString()
	c, err := pgxpool.Connect(ctx, ConnStr)
	if err != nil {
		log.Info().Msg("Zeus: InitPG failed to connect to database")
		panic(err)
	}

	Pg.DbMap[name] = c
	return c
}

func (d *Db) QueryRowWArgs(ctx context.Context, query string, args ...interface{}) pgx.Row {
	dbNameKey := ctx.Value("altDB")
	c := Pg.Pgpool
	if dbNameKey != nil {
		dbNameKeyStr := dbNameKey.(string)
		if len(dbNameKeyStr) > 0 {
			altC, ok := d.DbMap[dbNameKeyStr]
			if ok && altC != nil {
				c = altC
			}
		}
	}
	return c.QueryRow(ctx, query, args...)
}

func (d *Db) QueryRow(ctx context.Context, query string) pgx.Row {
	return Pg.Pgpool.QueryRow(ctx, query)
}

func (d *Db) QueryRow2(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return Pg.Pgpool.QueryRow(ctx, query, args...)
}

func (d *Db) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	return Pg.Pgpool.Query(ctx, query, args...)
}

func (d *Db) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return Pg.Pgpool.Exec(ctx, query, args...)
}

func (d *Db) Begin(ctx context.Context) (pgx.Tx, error) {
	return Pg.Pgpool.Begin(ctx)
}
