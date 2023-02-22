package main

import (
	"github.com/Soni295/shortlink/src/routes"
	"github.com/labstack/echo"
)

const PORT = ":8080"

func main() {
	e := echo.New()
	r := routes.Routers{}
	r.Build(e)
	e.Logger.Fatal(e.Start(PORT))
}
