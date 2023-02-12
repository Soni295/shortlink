package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

const PORT = ":8080"

func main() {
	fmt.Println("Server running on", PORT)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(PORT))
}
