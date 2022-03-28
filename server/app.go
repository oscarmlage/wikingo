package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/oscarmlage/wikingo/model"

	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

func WikiHome(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/WikiHome")
}

func WikiPage(c echo.Context) error {
	var page model.Page
	var err error
	page_id := c.Param("page")
	version := c.Param("version")
	if version != "" {
		Debug.Printf("version: %s", version)
		page, err = store.GetPageVersion(page_id, version)
		if err != nil {
			// If record not found, show 404
			if err.Error() == "record not found" {
				Debug.Printf("Page not found: %s/%s", page_id, version)
				return c.Render(http.StatusNotFound, "404.html", map[string]interface{}{
					"name": page_id,
				})
			}
		}
	} else {
		var err error
		page, err = store.GetPage(page_id)
		if err != nil {
			// If record not found, show the create
			if err.Error() == "record not found" {
				Debug.Printf("Page not found: %s", page_id)
				return c.Render(http.StatusOK, "notfound.html", map[string]interface{}{
					"name": page_id,
				})
			}
		}
	}
	Debug.Printf("res: %T\n", page)
	body := strings.ReplaceAll(page.Body, "\r\n", "\n")
	unsafe := blackfriday.Run([]byte(body))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	return c.Render(http.StatusOK, "page.html", map[string]interface{}{
		"page": page,
		"html": string(html[:]),
	})
}

func WikiPageEdit(c echo.Context) error {
	var page model.Page
	var err error
	page_id := c.Param("page")
	version := c.Param("version")
	if version != "" {
		Debug.Printf("version; %s\n", version)
		page, err = store.GetPageVersion(page_id, version)
	} else {
		Debug.Println("No version")
		page, err = store.GetPage(page_id)
	}
	if err != nil {
		// If record not found, create and show the form
		if err.Error() == "record not found" {
			err := store.AddPage(page_id, "# "+page_id)
			if err != nil {
				log.Panicln(err.Error())
			}
			Debug.Printf("Page %s created", page_id)
			page, err = store.GetPage(page_id)
			if err != nil {
				log.Panicln(err)
			}
		}
	}
	Debug.Printf("res: %T\n", page)
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
	Debug.Printf("res: %T\n", pages)
	return c.Render(http.StatusOK, "list.html", map[string]interface{}{
		"pages": pages,
	})
}

func WikiAbout(c echo.Context) error {
	return c.Render(http.StatusOK, "about.html", map[string]interface{}{})
}
