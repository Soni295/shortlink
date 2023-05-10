package drivens

import (
	"fmt"

	"github.com/Soni295/shortlink/src/services/user/app/entity"
	"github.com/Soni295/shortlink/src/services/user/ports/drivens"
)

func NewRepoQuerierStubAdapter() drivens.ForUserRepoQuery {
	return &RepoQuerierStubAdapter{}
}

var _ drivens.ForUserRepoQuery = &RepoQuerierStubAdapter{}

type RepoQuerierStubAdapter struct {
	users []entity.User
}

func (r *RepoQuerierStubAdapter) Save(user *entity.User) error {
	for _, userDb := range r.users {
		if user.Email == userDb.Email && user.Name == userDb.Name {
			return fmt.Errorf(ERR_USER_REGISTERED)
		}
	}

	r.users = append(r.users, *user)
	return nil
}

func (r *RepoQuerierStubAdapter) GetByMail(user *entity.User) (*entity.User, error) {
	for _, userDb := range r.users {
		if userDb.Email != user.Email {
			continue
		}
		return &userDb, nil
	}
	return nil, fmt.Errorf(ERR_USER_NOT_FOUND)
}
