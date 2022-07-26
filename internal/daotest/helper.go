//go:build integration_tests

package extensiontest

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/reymom/NFT-psql-go-interface/internal/dao"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

type LazyErrorHandler struct {
	t *testing.T
}

func (f LazyErrorHandler) HandleError(err error) {
	assert.NoError(f.t, err)
	return
}

type testEnvironment struct {
	*dao.PsqlDao
	ConnectionStringWrite string
}

func newTestEnvironment(readConnectionString string, writeConnectionString string, t *testing.T) (*testEnvironment, error) {
	psqlDao, e := dao.NewPsqlDao(
		&dao.Config{
			ConnectionStringRead:  readConnectionString,
			ConnectionStringWrite: writeConnectionString,
			WriteEnabled:          true,
			MaxReadConnections:    4,
			MaxWriteConnections:   4,
			AsyncErrorHandler:     LazyErrorHandler{t},
		},
	)
	if e != nil {
		return nil, e
	}
	return &testEnvironment{psqlDao, writeConnectionString}, nil
}

func (t *testEnvironment) cleanup() {
	const query = "TRUNCATE Users, Nfts " +
		"RESTART IDENTITY CASCADE"

	connection, e := pgx.Connect(context.Background(), t.ConnectionStringWrite)
	if e != nil {
		log.Err(e).Msg("")
		return
	}
	_, e = connection.Exec(context.Background(), query)
	if e != nil {
		log.Err(e).Msg("")
		return
	}
	return
}
