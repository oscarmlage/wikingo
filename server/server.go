package server

import (
	"log"

	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/oscarmlage/wikingo/model"
)

// Define the template registry struct
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

var (
	store model.Store
)

func Serve() {

	// Store instance
	// Depending on config we should open one store or other
	// (Gorm, File, Git...)
	store = new(model.StoreGorm)
	err := store.Open()
	if err != nil {
		log.Panicln(err)
	}

	// Echo instance
	e := echo.New()

	// Instantiate a template registry with an array of template set
	// Ref: https://gist.github.com/rand99/808e6e9702c00ce64803d94abff65678
	templates := make(map[string]*template.Template)

	templates["page.html"] = template.Must(template.New("page.html").Funcs(template.FuncMap{
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
	}).ParseFiles("views/page.html", "views/base.html"))
	templates["list.html"] = template.Must(template.ParseFiles("views/list.html", "views/base.html"))
	templates["edit.html"] = template.Must(template.ParseFiles("views/edit.html", "views/base.html"))
	templates["about.html"] = template.Must(template.ParseFiles("views/about.html", "views/base.html"))
	templates["notfound.html"] = template.Must(template.ParseFiles("views/notfound.html", "views/base.html"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	// Routes
	e.GET("/", WikiHome)
	e.GET("/:page", WikiPage)
	e.GET("/:page/edit", WikiPageEdit)
	e.POST("/:page/edit", WikiPagePostEdit)
	e.GET("/list", WikiList)
	e.GET("/about", WikiAbout)
	e.GET("/about/:id", WikiAbout)

	// Logger
	e.Logger.Fatal(e.Start(":2323"))
}
