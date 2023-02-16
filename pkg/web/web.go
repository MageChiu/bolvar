package web

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	//"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func startEchoServer(port int) {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	e.GET("/hello", Hello)
	adminGroup := e.Group("/admin",
		middleware.BasicAuth(func(username string,
			password string,
			context echo.Context) (bool, error) {
			e.Logger.Infof("username: %s", username)
			return true, nil
		}))
	//adminGroup.GET("/", func(c echo.Context) error {
	//	return c.HTML(http.StatusOK, "/static/admin-login.html")
	//})
	adminGroup.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"path": " test",
		})
	})
	//e.GET("/admin/infos", func(c echo.Context) error {
	//	return c.String(http.StatusOK, "test ok")
	//})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello.html", "World")
}
