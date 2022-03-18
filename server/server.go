package server

import (
	"github.com/labstack/echo/v4"
	"github.com/oscarmlage/wikingo/model"
	"log"
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
