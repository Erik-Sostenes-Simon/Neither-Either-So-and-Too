package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Proyecto/proyecto_ingles/database"
	"github.com/Proyecto/proyecto_ingles/routes"
	"github.com/labstack/echo/v4"
)

type Template struct {
	templates map[string]*template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]

	if !ok {
		err := errors.New("Error de página" + name)
		return fmt.Errorf("Error: %v", err)
	}
	return tmpl.ExecuteTemplate(w, "base.gohtml", data)
}
func main() {
	e := echo.New()

	templates := make(map[string]*template.Template)

	templates["home.gohtml"] = template.Must(template.ParseFiles("views/home.gohtml", "views/base.gohtml"))
	templates["form.gohtml"] = template.Must(template.ParseFiles("views/form.gohtml", "views/base.gohtml"))
	templates["answers.gohtml"] = template.Must(template.ParseFiles("views/answers.gohtml", "views/base.gohtml"))
	templates["list.gohtml"] = template.Must(template.ParseFiles("views/list.gohtml", "views/base.gohtml"))

	e.Renderer = &Template{
		templates: templates,
	}
	routes.Routes(e)
	db := database.ConexionDB()
	if err := db.Ping(); err != nil {
		panic(err)
	}
	e.Logger.Fatal(e.Start(":1323"))
}
