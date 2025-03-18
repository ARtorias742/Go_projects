package framework

import (
	"log"
	"net/http"

	"github.com/ARtorias742/GoForge/config"
	"github.com/ARtorias742/GoForge/watcher"
)

type App struct {
	Router     *Router
	Renderer   *TemplateRenderer
	Middleware Middleware
	Config     config.Config
}

func NewApp(cfg config.Config) *App {
	router := NewRouter()
	renderer, err := NewTemplateRenderer("templates")
	if err != nil {
		panic(err)
	}
	return &App{
		Router:     router,
		Renderer:   renderer,
		Middleware: NewMiddleware(),
		Config:     cfg,
	}
}

func (a *App) Start() {
	if a.Config.DebugMode {
		go watcher.StartWatcher("templates", "static", a.Config.Port)
	}
	log.Fatal(http.ListenAndServe(a.Config.Port, a.Router.Serve()))
}
