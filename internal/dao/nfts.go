package dao

import (
	"context"

	"github.com/reymom/NFT-psql-go-interface/pkg/model"
)

func (p *PsqlDao) CreateNFT(ctx context.Context, nft model.NFT) (*model.WrappedNFT, error) {
	panic("implement me")
}

func (p *PsqlDao) RemoveNFT(ctx context.Context, nftId string) error {
	panic("implement me")
}

func (p *PsqlDao) GetNFTById(ctx context.Context, nftId string) (*model.WrappedNFT, error) {
	panic("implement me")
}

func (p *PsqlDao) GetNFTsByOwnerId(ctx context.Context, ownerId string) ([]model.WrappedNFT, error) {
	panic("implement me")
}
