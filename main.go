package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/oscarmlage/wikingo/server"
	"html/template"
	"io"
)

// Define the template registry struct
type TemplateRegistry struct {
	templates *template.Template
}

// Implement e.Renderer intercace
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	fmt.Println("Starting wikingo...")

	server.Serve()
}
