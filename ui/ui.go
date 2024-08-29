package ui

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/mayowa/templates"
)

// wrapper struct to satisfy echo's renderer interface
type TemplateRenderer struct {
	AppTemplates *templates.Template
}

func NewTemplateRenderer() *TemplateRenderer {
	//TODO: REPLACE with path fig config
	tmpl, err := templates.New("./ui/templates", &templates.TemplateOptions{
		Ext: ".tmpl",
		FuncMap: template.FuncMap{"sub": func(a, b int) int {
			return a - b
		}},
		PathToSVG: "./assets/svg",
	})

	if err != nil {
		fmt.Println("error creating template renderer ", err)
		panic(err)
	}

	renderer := TemplateRenderer{AppTemplates: tmpl}
	return &renderer
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.AppTemplates.Render(w, name, data)
}
