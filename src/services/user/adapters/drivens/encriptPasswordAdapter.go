package drivens

import (
	"fmt"

	"github.com/Soni295/shortlink/src/services/user/ports/drivens"
	"golang.org/x/crypto/bcrypt"
)

func NewEncryptPasswordAdapter() drivens.ForEncryptPassword {
	return &EncryptPasswordAdapter{}
}

var _ drivens.ForEncryptPassword = &EncryptPasswordAdapter{}

type EncryptPasswordAdapter struct{}

func (e *EncryptPasswordAdapter) Encript(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (e *EncryptPasswordAdapter) CheckPassword(password string, passwordEncrypted string) error {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(passwordEncrypted),
		[]byte(password),
	); err != nil {
		return fmt.Errorf(ERR_PASSWORD_INCORRECT)
	}
	return nil
}
