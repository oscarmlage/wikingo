package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func WikiHome(c echo.Context) error {
	res, err := store.GetAllPages()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("res: %T", res)
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name":  "HOME",
		"msg":   "This is the home page.",
		"pages": res,
	})
}

func WikiAbout(c echo.Context) error {
	id := c.Param("id")
	return c.Render(http.StatusOK, "about.html", map[string]interface{}{
		"name": "ABOUT",
		"msg":  "About Wikingo (id:)" + id,
	})
}
