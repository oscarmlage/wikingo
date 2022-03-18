package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func WikiHome(c echo.Context) error {
	res := store.GetPage()
	fmt.Println(res)
	return c.String(http.StatusOK, "WikiHome")
}

func WikiAbout(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "About the wiki. id:"+id)
}

func WikiListPages(c echo.Context) error {
	fmt.Println("WikiListPages")
	res, err := store.GetAllPages()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(res)
	return c.String(http.StatusOK, "WikiListPages")
}
