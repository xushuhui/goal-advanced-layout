package biz

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	v1 "goal-advanced-layout/api"
	"goal-advanced-layout/internal/data/model"
)

type UserRepo interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}
type UserUsecase struct {
	c  *Usecase
	ur UserRepo
}

func NewUserUsecase(c *Usecase, userRepo UserRepo) *UserUsecase {
	return &UserUsecase{
		ur: userRepo,
		c:  c,
	}
}

func (s *UserUsecase) Register(ctx context.Context, req *v1.RegisterRequest) error {
	// check username
	if user, err := s.ur.GetByUsername(ctx, req.Username); err == nil && user != nil {
		return v1.ErrUsernameAlreadyUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// Generate user ID
	userID, err := s.c.sid.GenString()
	if err != nil {
		return err
	}
	// Create a user
	user := &model.User{
		UserId:   userID,
		Username: req.Username,
		Nickname: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
	}
	if err = s.ur.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *UserUsecase) Login(ctx context.Context, req *v1.LoginRequest) (string, error) {
	user, err := s.ur.GetByUsername(ctx, req.Username)
	if err != nil || user == nil {
		return "", v1.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}
	token, err := s.c.jwt.GenToken(user.UserId, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserUsecase) GetProfile(ctx context.Context, userId string) (*v1.GetProfileResponseData, error) {
	user, err := s.ur.GetByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1.GetProfileResponseData{
		UserId:   user.UserId,
		Nickname: user.Nickname,
		Username: user.Username,
	}, nil
}

func (s *UserUsecase) UpdateProfile(ctx context.Context, userId string, req *v1.UpdateProfileRequest) error {
	user, err := s.ur.GetByID(ctx, userId)
	if err != nil {
		return err
	}

	user.Email = req.Email
	user.Nickname = req.Nickname

	if err = s.ur.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
