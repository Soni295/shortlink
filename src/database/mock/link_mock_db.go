package mock

import (
	"fmt"

	"github.com/Soni295/shortlink/src/database"
	"github.com/Soni295/shortlink/src/entity"
)

type LinkMockDb struct {
	links []entity.Link
}

func NewLinkDbMock() (database.ILinkDb, error) {
	return &LinkMockDb{[]entity.Link{
		{
			Url:     "http://www.youtube.com",
			Shorted: "yutu",
		},
	}}, nil
}

func (lmdb *LinkMockDb) GET(shortName string) (*entity.Link, error) {
	for _, item := range lmdb.links {
		if item.Shorted == shortName {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("ShortLink Not found")
}

func (lmdb *LinkMockDb) Add(link *entity.Link) (*entity.Link, error) {
	lmdb.links = append(lmdb.links, *link)
	return link, nil
}
