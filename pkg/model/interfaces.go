package model

import "context"

type NFTStoreDao interface {
	CreateNFT(ctx context.Context, nft NFT) (*WrappedNFT, error)
	RemoveNFT(ctx context.Context, nftId string) error
	GetNFTById(ctx context.Context, nftId string) (*WrappedNFT, error)
	GetNFTsByOwnerId(ctx context.Context, ownerId string) ([]WrappedNFT, error)
}

type UserDao interface {
	AddUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, userId string) error
	GetUsers(ctx context.Context) ([]User, error)
	GetUserByUserId(ctx context.Context, userId string) (*User, error)
}

type NFT interface {
	GetNFTTypeId() NFTTypeId
	ToJson() ([]byte, error)
}
