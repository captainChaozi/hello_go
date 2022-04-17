package hello_go

import (
	"github.com/labstack/echo"
	"net/http"
)

const A = 10

func GoWeb() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
