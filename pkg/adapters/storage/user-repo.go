package storage

import (
	"context"
	"errors"
	"tasoryx/internal/user/domain"
	"tasoryx/internal/user/port"
	"tasoryx/pkg/adapters/storage/mapper"
	"tasoryx/pkg/adapters/storage/types"
	"tasoryx/pkg/fp"
	"time"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) port.Repo {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) Create(ctx context.Context, data domain.User) (domain.UserID, error) {
	user := mapper.UserDomain2Repo(data)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	result := ur.db.WithContext(ctx).Create(user)
	if result.Error != nil {
		return user.ID, errors.New("failed to create user : " + result.Error.Error())
	}
	return user.ID, nil
}
func (ur *userRepo) GetByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	user := new(types.User)
	result := ur.db.WithContext(ctx).First(user, "id = ?", userID)
	userDomain := mapper.UserRepo2Domain(*user)
	if result.Error != nil {
		return userDomain, errors.New("failed to get user by id : " + result.Error.Error())
	}
	return userDomain, nil
}
func (ur *userRepo) Get(ctx context.Context, filter ...domain.FilterUser) ([]domain.User, error) {
	var users []types.User
	result := ur.db.WithContext(ctx).Find(&users, filter)
	if result.Error != nil {
		return nil, errors.New("failed to get users : " + result.Error.Error())
	}
	usersDomain := fp.Map(users, mapper.UserRepo2Domain)
	return usersDomain, nil
}
func (ur *userRepo) Update(ctx context.Context, NewRecord domain.User, ID domain.UserID) error {
	user, err := ur.GetByID(ctx, ID)
	if err != nil {
		return errors.New("failed to get user by id: " + err.Error())
	}
	newUser := mapper.UserDomain2Repo(NewRecord)
	newUser.UpdatedAt = time.Now()
	result := ur.db.WithContext(ctx).Model(user).Updates(newUser)
	if result.Error != nil {
		return errors.New("failed to update user: " + result.Error.Error())
	}
	return nil
}
func (ur *userRepo) Delete(ctx context.Context, ID domain.UserID) error {
	result := ur.db.WithContext(ctx).Delete(&types.User{}, ID)
	if result.Error != nil {
		return errors.New("failed to update user: " + result.Error.Error())
	}
	return nil
}
