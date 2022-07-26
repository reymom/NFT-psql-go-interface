package dao

import (
	"context"

	"github.com/reymom/NFT-psql-go-interface/pkg/model"
)

func (p *PsqlDao) AddUser(ctx context.Context, user *model.User) error {
	panic("implement me")
}

func (p *PsqlDao) DeleteUser(ctx context.Context, userId string) error {
	panic("implement me")
}

func (p *PsqlDao) GetUsers(ctx context.Context) ([]model.User, error) {
	panic("implement me")
}

func (p *PsqlDao) GetUserByUserId(ctx context.Context, userId string) (*model.User, error) {
	panic("implement me")
}
