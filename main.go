package main

import (
	"main/router"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	router.Init(e)
	e.Logger.Fatal(e.Start(":1323"))
}
