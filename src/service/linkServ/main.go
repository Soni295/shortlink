package linkserv

import (
	"fmt"

	"github.com/Soni295/shortlink/src/database"
	"github.com/Soni295/shortlink/src/entity"
	"github.com/Soni295/shortlink/src/utils"
)

type ILinkServ interface {
	CreateShortName(link *entity.Link) (*entity.Link, error)
	GetByShortName(shortName string) (*entity.Link, error)
}

type LinkServ struct {
	longShortNames int
	linkDb         database.ILinkDb
}

func New(linkdb database.ILinkDb, longShortNames int) (ILinkServ, error) {
	if linkdb == nil {
		return nil, fmt.Errorf(ErrDatabaseIsNotProvider)
	}

	if longShortNames < SHORTNAMELENGTH {
		return nil, fmt.Errorf(ErrLenghtShortNameIsTooShort)
	}

	return &LinkServ{
		longShortNames: longShortNames,
		linkDb:         linkdb,
	}, nil
}

func (lServ *LinkServ) CreateShortName(link *entity.Link) (*entity.Link, error) {
	link.Shorted = utils.GenerateShorNameHandler(lServ.longShortNames)
	complete, err := lServ.linkDb.Add(link)

	if err != nil {
		return nil, err
	}
	return complete, nil
}

func (lServ *LinkServ) GetByShortName(shortName string) (*entity.Link, error) {
	return lServ.linkDb.GET(shortName)
}
