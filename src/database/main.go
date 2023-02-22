package database

import "github.com/Soni295/shortlink/src/entity"

type ILinkDb interface {
	GET(shortName string) (*entity.Link, error)
	Add(link *entity.Link) (*entity.Link, error)
}
