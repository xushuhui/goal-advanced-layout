package data

import (
	"context"

	"goal-advanced-layout/internal/biz"
	"goal-advanced-layout/internal/data/model"
)

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		Data: data,
	}
}

type userRepo struct {
	*Data
}

func (r *userRepo) Create(ctx context.Context, user *biz.User) error {
	_, err := r.query.CreateUser(ctx, model.CreateUserParams{})
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) Update(ctx context.Context, user *biz.User) error {
	return nil
}

func (r *userRepo) GetByID(ctx context.Context, userID string) (*biz.User, error) {
	user, err := r.query.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		UserId: user.UserID,
	}, nil
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*biz.User, error) {
	user, err := r.query.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		UserId: user.UserID,
	}, nil
}
