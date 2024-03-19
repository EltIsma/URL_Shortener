package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Database struct {
	db *pgxpool.Pool
}

func NewDatabase(dbConf string) (*Database, error) {
	config, err := pgxpool.ParseConfig(dbConf)
	if err != nil {
		logrus.WithError(err).Fatalf("can't parse pgxpool config")
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("can't create pg pool: %s", err.Error())
	}
	conn, err := pool.Acquire(context.Background())
    if err != nil {
        return nil, fmt.Errorf("can't connect to database: %s", err.Error())
    }
    conn.Release()

	return &Database{db: pool}, nil
}

func (d *Database) GetDB() *pgxpool.Pool {
	return d.db
}
