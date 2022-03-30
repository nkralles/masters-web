package pgdriver

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nkralles/masters-web/internal/cfg"
	"github.com/nkralles/masters-web/internal/logger"
	"github.com/nkralles/masters-web/internal/persistence"
	"gitlab.cloud.n-ask.com/n-ask/fancylog"
	"net/url"
)

type Driver struct {
	pool *pgxpool.Pool
}

func Connect(db string) (persistence.MastersStorage, error) {
	var err error
	var dbURL *url.URL
	dbURL, err = url.Parse(db)
	if err != nil {
		logger.Internal.Fatalf("Invalid url: %v", err)
	}
	q := dbURL.Query()
	if len(q.Get("sslmode")) == 0 {
		q.Set("sslmode", "disable")
	}
	dbURL.RawQuery = q.Encode()

	//fmt.Println(dbURL)
	//fmt.Println(cfg.Config.DatabaseMigrationPath)
	//var m *migrate.Migrate
	//if m, err = migrate.New(cfg.Config.DatabaseMigrationPath, dbURL.String()); err != nil {
	//	log.Fatalf("Migration init failed: %v", err)
	//}
	//if err = m.Up(); err != nil && err != migrate.ErrNoChange {
	//	log.Fatalf("Migration failed: %v", err)
	//}
	connConfig, err := pgxpool.ParseConfig(dbURL.String())
	if err != nil {
		return nil, err
	}
	if cfg.Config.Verbose {
		connConfig.ConnConfig.Logger = fancylog.NewLogger(logger.PG)
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		return nil, err
	}

	d := &Driver{pool: pool}

	return d, nil
}
