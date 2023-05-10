package app

import (
	"github.com/Soni295/shortlink/src/services/user/app/entity"
	"github.com/Soni295/shortlink/src/services/user/ports/drivens"
	"github.com/Soni295/shortlink/src/services/user/ports/drivers"
)

type UserManager struct {
	handlerEncrypt drivens.ForEncryptPassword
	userRepo       drivens.ForUserRepoQuery
}

var _ drivers.ForManagingUser = &UserManager{}

func NewUserManager(handlerEncrypt drivens.ForEncryptPassword, userRepo drivens.ForUserRepoQuery) *UserManager {
	return &UserManager{
		handlerEncrypt: handlerEncrypt,
		userRepo:       userRepo,
	}
}

func (u *UserManager) SaveUser(name string, email string, password string) (*entity.UserWithoutPassword, error) {
	hashPassword, err := u.handlerEncrypt.Encript(password)
	if err != nil {
		return nil, err
	}

	user := entity.User{
		Email:    email,
		Name:     name,
		Password: hashPassword,
	}

	if err := u.userRepo.Save(&user); err != nil {
		return nil, err
	}

	return &entity.UserWithoutPassword{
		Name:  name,
		Email: email,
	}, nil
}

func (u *UserManager) GetUserByEmailAndPassword(email string, password string) (*entity.UserWithoutPassword, error) {
	user, err := u.userRepo.GetByMail(&entity.User{Email: email})
	if err != nil {
		return nil, err
	}
	if err := u.handlerEncrypt.CheckPassword(password, user.Password); err != nil {
		return nil, err
	}

	return &entity.UserWithoutPassword{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
