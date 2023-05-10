package user_test

import (
	"testing"

	"github.com/Soni295/shortlink/src/services/user"
	"github.com/Soni295/shortlink/src/services/user/adapters/drivens"
	"github.com/Soni295/shortlink/src/services/user/app/entity"
	"github.com/stretchr/testify/assert"
)

func TestUserManager(t *testing.T) {
	userManagerMock := user.CompositionMock()

	t.Run("Sign up user success", func(t *testing.T) {

		userExpected := &entity.UserWithoutPassword{
			Name:  "martin",
			Email: "martin@gmail.com",
		}

		t.Run("success", func(t *testing.T) {
			user, err := userManagerMock.SaveUser(userExpected.Name, userExpected.Email, "123456")
			assert.Nil(t, err)
			assert.Equal(t, userExpected, user)
		})

		t.Run("fail because email duplicated", func(t *testing.T) {
			user, err := userManagerMock.SaveUser(userExpected.Name, userExpected.Email, "123456")
			assert.Nil(t, user)
			assert.EqualValues(t, err.Error(), drivens.ERR_USER_REGISTERED)
		})
	})

	t.Run("Get user by email and password", func(t *testing.T) {
		userReference := &entity.User{
			Name:     "joe",
			Email:    "joedoe@gmail.com",
			Password: "123456",
		}

		userExpected := &entity.UserWithoutPassword{
			Name:  userReference.Name,
			Email: userReference.Email,
		}

		userOutput, err := userManagerMock.SaveUser(userReference.Name, userReference.Email, userReference.Password)
		assert.Nil(t, err)
		assert.Equal(t, userExpected, userOutput)

		t.Run("success", func(t *testing.T) {
			userOutput, err := userManagerMock.GetUserByEmailAndPassword(userReference.Email, userReference.Password)
			assert.Nil(t, err)
			assert.Equal(t, userExpected, userOutput)
		})

		t.Run("fail because the password is wrong", func(t *testing.T) {
			userNoExpected, err := userManagerMock.GetUserByEmailAndPassword(userReference.Email, "wrong passoword")
			assert.Nil(t, userNoExpected)
			assert.Equal(t, err.Error(), drivens.ERR_PASSWORD_INCORRECT)
		})

		t.Run("fail because user doesn´t exist", func(t *testing.T) {
			userNoExpected, err := userManagerMock.GetUserByEmailAndPassword("example@gmail.com", "wrong passoword")
			assert.Nil(t, userNoExpected)
			assert.Equal(t, err.Error(), drivens.ERR_USER_NOT_FOUND)
		})
	})
}
