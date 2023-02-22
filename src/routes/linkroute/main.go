package linkroute

import (
	"net/http"

	"github.com/Soni295/shortlink/src/controllers/linkctrl"
	"github.com/Soni295/shortlink/src/entity"
	"github.com/labstack/echo"
)

func New(linkctrl linkctrl.ILinkCtrl) LinkRoute {
	return LinkRoute{linkctrl}
}

type LinkRoute struct {
	linkctrl linkctrl.ILinkCtrl
}

func (lr *LinkRoute) GetByShortName(c echo.Context) error {
	shortName := c.Param("shortName")

	link, err := lr.linkctrl.GetByShortName(shortName)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	return c.Redirect(301, link.Url)
}

func (lr *LinkRoute) CreateShortName(c echo.Context) error {
	link := &entity.Link{}
	if err := c.Bind(link); err != nil {
		return err
	}
	newLink, err := lr.linkctrl.CreateShortName(link)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusCreated, newLink)
}
