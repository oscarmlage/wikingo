package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func WikiHome(c echo.Context) error {
	var err error
	page, err := store.GetPage("WikiHome")
	if err != nil {
		fmt.Println(err.Error())
		err := store.AddPage("WikiHome", "This is the home of the wiki")
		if err != nil {
			log.Panicln(err.Error())
		}
	}
	fmt.Printf("res: %T\n", page)
	return c.Render(http.StatusOK, "page.html", map[string]interface{}{
		"page": page,
	})
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
	return c.Render(http.StatusOK, "page.html", map[string]interface{}{
		"page": page,
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
