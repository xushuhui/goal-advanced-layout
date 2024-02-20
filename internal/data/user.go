package data

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"goal-advanced-layout/api"
	"goal-advanced-layout/internal/data/model"
)

type UserRepo interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}

func NewUserRepo(data *Data) UserRepo {
	return &userRepo{
		Data: data,
	}
}

type userRepo struct {
	*Data
}

func (r *userRepo) Create(ctx context.Context, user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepo) Update(ctx context.Context, user *model.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepo) GetByID(ctx context.Context, userId string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, api.ErrNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
