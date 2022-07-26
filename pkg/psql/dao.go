package psql

import (
	"github.com/reymom/NFT-psql-go-interface/internal/nft"
	"github.com/reymom/NFT-psql-go-interface/pkg/model"
)

type ConfigPsql struct {
	ConnectionStringRead  string
	ConnectionStringWrite string
	MaxReadConnections    uint
	MaxWriteConnections   uint
	AsyncErrorHandler     model.ErrorHandler
}

func NewNFTStoreDao(config *ConfigPsql) (model.NFTStoreDao, error) {
	return nft.NewPsqlDao(&nft.Config{
		ConnectionStringRead:  config.ConnectionStringRead,
		ConnectionStringWrite: config.ConnectionStringWrite,
		WriteEnabled:          true,
		MaxReadConnections:    config.MaxReadConnections,
		MaxWriteConnections:   config.MaxWriteConnections,
		AsyncErrorHandler:     config.AsyncErrorHandler,
	})
}

func NewUserStoreDao(config *ConfigPsql) (model.UserDao, error) {
	return nft.NewPsqlDao(&nft.Config{
		ConnectionStringRead:  config.ConnectionStringRead,
		ConnectionStringWrite: config.ConnectionStringWrite,
		WriteEnabled:          true,
		MaxReadConnections:    config.MaxReadConnections,
		MaxWriteConnections:   config.MaxWriteConnections,
		AsyncErrorHandler:     config.AsyncErrorHandler,
	})
}
