package drivers

import "github.com/Soni295/shortlink/src/services/user/app/entity"

type ForManagingUser interface {
	GetUserByEmailAndPassword(email, password string) (*entity.UserWithoutPassword, error)
	SaveUser(name, email, password string) (*entity.UserWithoutPassword, error)
}
