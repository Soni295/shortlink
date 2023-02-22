package linkctrl

import (
	"github.com/Soni295/shortlink/src/entity"
	linkserv "github.com/Soni295/shortlink/src/service/linkServ"
)

type ILinkCtrl interface {
	GetByShortName(shortName string) (*entity.Link, error)
	CreateShortName(link *entity.Link) (*entity.Link, error)
}

type LinkCtrl struct {
	linkServ linkserv.ILinkServ
}

func New(linkserv linkserv.ILinkServ) ILinkCtrl {
	return &LinkCtrl{
		linkServ: linkserv,
	}
}

func (lc *LinkCtrl) GetByShortName(shortName string) (*entity.Link, error) {
	return lc.linkServ.GetByShortName(shortName)
}

func (lc *LinkCtrl) CreateShortName(link *entity.Link) (*entity.Link, error) {
	return lc.linkServ.CreateShortName(link)
}
