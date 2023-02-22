package routes

import (
	"log"

	"github.com/Soni295/shortlink/src/controllers/linkctrl"
	"github.com/Soni295/shortlink/src/infrastucture/persistence"
	"github.com/Soni295/shortlink/src/routes/linkroute"
	linkserv "github.com/Soni295/shortlink/src/service/linkServ"
	"github.com/labstack/echo"
)

type Routers struct{}

func (r *Routers) Build(e *echo.Echo) *echo.Echo {
	repo, err := persistence.NewRepositories()
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	serv, err := linkserv.New(repo.Link, 10)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	ctrl := linkctrl.New(serv)
	link := linkroute.New(ctrl)

	linkR := e.Group("/link")
	linkR.GET("/:shortName", link.GetByShortName)
	linkR.POST("", link.CreateShortName)
	return e
}
