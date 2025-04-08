package components

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Application) Mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})
	return r
}
func (a *Application) Run(mux *chi.Mux) error {
	srv := &http.Server{
		Addr:         a.Config.Addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Сервер запущен на %s", a.Config.Addr)
	return srv.ListenAndServe()
}
func (a *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK: healthCheckHandler сработал"))
	// ? will think about it:
	// !a.Store.Posts.Create(r.Context())
}
