package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/oscarmlage/wikingo/model"
	"log"
	"net/http"
)

// Depending on config we should open one store or other (Gorm, File,
// Git...)
var (
	store model.StoreGorm
)

func Serve() {
	// Store instance
	err := store.Open()
	if err != nil {
		log.Panicln(err)
	}

	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", WikiHome)
	e.GET("/list", WikiListPages)
	e.GET("/about", WikiAbout)
	e.GET("/about/:id", WikiAbout)

	// Logger
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

func WikiListPages(c echo.Context) error {
	fmt.Println("WikiListPages")
	res, err := store.GetAllPages()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(res)
	return c.String(http.StatusOK, "WikiListPages")
}
