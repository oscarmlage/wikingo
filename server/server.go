package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/oscarmlage/wikingo/model"
	"log"
	"net/http"
)

var store *model.Store

func Serve() {
	store = &model.Store{}
	err := store.Open()
	if err != nil {
		log.Panicln(err)
	}
	e := echo.New()
	e.GET("/", WikiHome)
	e.GET("/about", WikiAbout)
	e.GET("/about/:id", WikiAbout)

	e.Logger.Fatal(e.Start(":2323"))
}

func WikiHome(c echo.Context) error {
	res := store.GetPage()
	fmt.Println(res)
	return c.String(http.StatusOK, "WikiHome")
}

func WikiAbout(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "About the wiki. id:"+id)
}
