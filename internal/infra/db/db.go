package db

import (
	"PaymentProcessingSystem/internal"
	"PaymentProcessingSystem/internal/configs"

	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewDB instantiates the DB database using configuration defined in environment variables.
func NewDB(conf *configs.DatabaseConfig) (*pgxpool.Pool, error) {
	dsn := conf.BuildDSN()

	pool, err := pgxpool.New(context.Background(), dsn.String())
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "pgxpool.Connect")
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "db.Ping")
	}

	return pool, nil
}
