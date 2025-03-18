package main

import (
	"log"
	"net/http"

	"github.com/ARtorias742/GoForge/config"
	"github.com/ARtorias742/GoForge/framework"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize framework
	app := framework.NewApp(cfg)

	// Serve static files
	app.Router.ServeStatic("/static/", "static")

	// Define routes
	app.Router.Handle("/", app.Middleware.Logging(func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"Title":   "AdvGoFront Home",
			"Message": "Welcome to the advanced Go frontend!",
		}
		app.Renderer.Render(w, "index.html", data)
	}))

	// Form handling example
	app.Router.Handle("/submit", app.Middleware.Logging(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			formData := framework.ParseForm(r)
			app.Renderer.Render(w, "index.html", map[string]interface{}{
				"Title":   "Form Submitted",
				"Message": "Received: " + formData["name"],
			})
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}))

	// Start server with live reloading
	log.Printf("Starting server on %s", cfg.Port)
	app.Start()
}
