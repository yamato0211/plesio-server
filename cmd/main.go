package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	num, err := strconv.Atoi("1")
	if err != nil {
		fmt.Println("hoge")
	}
	fmt.Println(num)

	num, err = strconv.Atoi("2")
	fmt.Println(num)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})
	e.Logger.Fatal(e.Start(":8000"))
}
