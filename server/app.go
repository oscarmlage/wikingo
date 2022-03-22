package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"log"
	"net/http"
)

func WikiHome(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/WikiHome")
}

func WikiPage(c echo.Context) error {
	page_id := c.Param("page")
	var err error
	page, err := store.GetPage(page_id)
	if err != nil {
		// If record not found, show the create
		if err.Error() == "record not found" {
			return c.Render(http.StatusOK, "notfound.html", map[string]interface{}{
				"name": page_id,
			})
		}
	}
	fmt.Printf("res: %T\n", page)
	body := strings.ReplaceAll(page.Body, "\r\n", "\n")
	unsafe := blackfriday.Run([]byte(body))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return c.Render(http.StatusOK, "page.html", map[string]interface{}{
		"page": page,
		"html": string(html[:]),
	})
}

func WikiPageEdit(c echo.Context) error {
	page_id := c.Param("page")
	var err error
	page, err := store.GetPage(page_id)
	if err != nil {
		// If record not found, create and show the form
		if err.Error() == "record not found" {
			err := store.AddPage(page_id, "# "+page_id)
			if err != nil {
				log.Panicln(err.Error())
			}
			page, err = store.GetPage(page_id)
			if err != nil {
				log.Panicln(err)
			}
		}
	}
	fmt.Printf("res: %T\n", page)
	return c.Render(http.StatusOK, "edit.html", map[string]interface{}{
		"page": page,
	})
}

func WikiPagePostEdit(c echo.Context) error {
	page_id := c.Param("page")
	var err error
	page, err := store.GetPage(page_id)
	if err != nil {
		log.Panic(err.Error())
	}
	page.Body = c.FormValue("html")
	store.CreatePage(page)
	return c.Redirect(http.StatusMovedPermanently, "/"+page.Name)
}

func WikiList(c echo.Context) error {
	pages, err := store.GetAllPages()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("res: %T\n", pages)
	return c.Render(http.StatusOK, "list.html", map[string]interface{}{
		"pages": pages,
	})
}

func WikiAbout(c echo.Context) error {
	return c.Render(http.StatusOK, "about.html", map[string]interface{}{})
}
