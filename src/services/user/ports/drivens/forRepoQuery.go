package drivens

import "github.com/Soni295/shortlink/src/services/user/app/entity"

type ForUserRepoQuery interface {
	GetByMail(user *entity.User) (*entity.User, error)
	Save(user *entity.User) error
}
