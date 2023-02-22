package linkctrl_test

import (
	"testing"

	"github.com/Soni295/shortlink/src/database/mock"
	linkserv "github.com/Soni295/shortlink/src/service/linkServ"
	"github.com/stretchr/testify/assert"
)

func TestLinkCtrl(t *testing.T) {

	t.Run("New Struct", func(t *testing.T) {
		t.Run("It should returns (nil and error)", func(t *testing.T) {
			linkS, err := linkserv.New(nil, 0)

			assert.Nil(t, linkS)
			assert.Error(t, err)
		})

		t.Run("It should returns (nil and error)", func(t *testing.T) {
			linkS, err := linkserv.New(nil, 0)
			errExpected := linkserv.ErrDatabaseIsNotProvider

			assert.Nil(t, linkS)
			assert.EqualError(t, err, errExpected)
		})

		t.Run("It should return (nil and error)", func(t *testing.T) {
			mockDb := mock.NewLinkDbMock()
			linkS, err := linkserv.New(mockDb, 0)
			errExpected := linkserv.ErrLenghtShortNameIsTooShort

			assert.Nil(t, linkS)
			assert.Error(t, err)
			assert.EqualError(t, err, errExpected)
		})

		t.Run("It should return (nil and error)", func(t *testing.T) {
			mockDb := mock.NewLinkDbMock()
			linkS, err := linkserv.New(mockDb, 3)
			errExpected := linkserv.ErrLenghtShortNameIsTooShort

			assert.Nil(t, linkS)
			assert.Error(t, err)
			assert.EqualError(t, err, errExpected)
		})

		t.Run("It should return (LinkServ and nil)", func(t *testing.T) {
			mockDb := mock.NewLinkDbMock()
			linkS, err := linkserv.New(mockDb, 4)
			assert.Nil(t, err)
			assert.NotNil(t, linkS)
		})
	})

}
