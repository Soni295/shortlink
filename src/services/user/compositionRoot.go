package user

import (
	"github.com/Soni295/shortlink/src/services/user/adapters/drivens"
	adapterdriver "github.com/Soni295/shortlink/src/services/user/adapters/drivers"
	"github.com/Soni295/shortlink/src/services/user/app"
	"github.com/Soni295/shortlink/src/services/user/ports/drivers"
)

func CompositionMock() drivers.ForManagingUser {
	userRepo := drivens.NewRepoQuerierStubAdapter()
	handlerEncrypt := drivens.NewEncryptPasswordAdapter()
	userManager := app.NewUserManager(handlerEncrypt, userRepo)
	userManagerProxy := adapterdriver.NewUserManagerProxy(userManager)
	return userManagerProxy
}
