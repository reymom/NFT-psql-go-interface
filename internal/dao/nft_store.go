package dao

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/reymom/NFT-psql-go-interface/pkg/model"
)

type Config struct {
	ConnectionStringRead  string
	ConnectionStringWrite string
	WriteEnabled          bool
	MaxReadConnections    uint
	MaxWriteConnections   uint
	AsyncErrorHandler     model.ErrorHandler
}

type PsqlDao struct {
	readPool          *pgxpool.Pool
	writePool         *pgxpool.Pool
	readEnabled       bool
	writeEnabled      bool
	asyncErrorHandler model.ErrorHandler
}

func (p *PsqlDao) handleAsyncError(e error) {
	if p.asyncErrorHandler != nil && e != nil {
		p.asyncErrorHandler.HandleError(e)
	}
}

func NewPsqlDao(config *Config) (*PsqlDao, error) {
	if config == nil {
		return nil, model.ErrNilNotAllowed
	}
	readPool, e := initializePool(config.ConnectionStringRead, config.MaxReadConnections)
	if e != nil {
		return nil, e
	}

	var writePool *pgxpool.Pool
	switch config.WriteEnabled {
	case true:
		writePool, e = initializePool(config.ConnectionStringWrite, config.MaxWriteConnections)
		if e != nil {
			return nil, e
		}
	case false:
		writePool, e = initializeEmptyPool(model.ErrFakeWritePool)
		if e != nil {
			return nil, e
		}
	}

	return &PsqlDao{
		readPool:          readPool,
		writePool:         writePool,
		readEnabled:       true,
		writeEnabled:      config.WriteEnabled,
		asyncErrorHandler: config.AsyncErrorHandler,
	}, nil
}

func initializePool(connectionString string, maxPoolSize uint) (*pgxpool.Pool, error) {

	poolConfig, e := pgxpool.ParseConfig(connectionString)
	if e != nil {
		return nil, e
	}
	poolConfig.MaxConns = int32(maxPoolSize)

	//We error after 11 seconds no connection can be made
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*11)
	defer cancel()

	return pgxpool.ConnectConfig(ctx, poolConfig)
}

func initializeEmptyPool(fakePoolError error) (*pgxpool.Pool, error) {
	const fakeConnectionString = "postgresql://no-user:password@nonhost:1/non-existing-db"
	poolConfig, e := pgxpool.ParseConfig(fakeConnectionString)
	if e != nil {
		return nil, e
	}
	poolConfig.LazyConnect = true
	poolConfig.BeforeConnect = func(ctx context.Context, config *pgx.ConnConfig) error {
		return fakePoolError
	}

	return pgxpool.ConnectConfig(context.Background(), poolConfig)
}
