package data

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"goal-advanced-layout/api"
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

func (r *userRepo) GetByID(ctx context.Context, userID string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
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
