package app

import (
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

// Template - чтобы добавить в Echo собственный шаблонизатор нужно создать структуру, которая реализует интерфейс echo.Renderer.
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, _ echo.Context) error {
	return errors.Wrap(t.templates.ExecuteTemplate(w, name, data), "can not execute templates")
}

func (s *Server) Start() {
	e := echo.New()

	// Добавим middleware которое будет выводить сообщения об ошибках в консоль
	e.Use(middleware.Logger())

	// Подключим наш шаблонизатор
	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("src/templates/*.html")),
	}

	//e.GET("/", handlers.Main(s.DB))
	//
	//e.GET("/product/:id/", handlers.Product(s.DB))

	if err := e.Start("localhost:8000"); err != nil {
		log.Fatal(err)
	}
}
