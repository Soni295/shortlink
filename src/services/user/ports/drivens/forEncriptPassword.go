package drivens

type ForEncryptPassword interface {
	Encript(password string) (string, error)
	CheckPassword(password, passwordEncrypted string) error
}
