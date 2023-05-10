package drivers

import (
	"github.com/Soni295/shortlink/src/services/user/app"
	"github.com/Soni295/shortlink/src/services/user/app/entity"
	"github.com/Soni295/shortlink/src/services/user/ports/drivers"
)

type UserManagerProxy struct {
	repository *app.UserManager
}

var _ drivers.ForManagingUser = &UserManagerProxy{}

func NewUserManagerProxy(repository *app.UserManager) drivers.ForManagingUser {
	return &UserManagerProxy{
		repository: repository,
	}
}

func (u *UserManagerProxy) SaveUser(name string, email string, password string) (*entity.UserWithoutPassword, error) {
	return u.repository.SaveUser(name, email, password)
}

func (u *UserManagerProxy) GetUserByEmailAndPassword(email string, password string) (*entity.UserWithoutPassword, error) {
	return u.repository.GetUserByEmailAndPassword(email, password)
}
