package components

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Mount настраивает маршруты HTTP
func (app *Application) Mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)                 // Уникальный ID для каждого запроса
	r.Use(middleware.RealIP)                    // Определение реального IP клиента
	r.Use(middleware.Logger)                    // Логирование запросов
	r.Use(middleware.Recoverer)                 // Восстановление после сбоев
	r.Use(middleware.Timeout(60 * time.Second)) // Ограничение времени запроса
	r.Route("/v1", func(r chi.Router) {
		// Проверка состояния сервера
		r.Get("/status", app.statusCheckHandler)
	})
	return r
}

// Run запускает сервер
func (app *Application) Run(mux *chi.Mux) error {
	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Printf("Сервер запущен на %s", app.Config.Addr)
	return srv.ListenAndServe()
}

// statusCheckHandler проверяет состояние сервера
// @Summary Проверка состояния сервера
// @Description Возвращает "OK", если сервер работает
// @Tags health
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /v1/status [get]
func (app *Application) statusCheckHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("OK")); err != nil {
		log.Printf("Ошибка ответа: %v", err)
		http.Error(w, "Ошибка сервера", http.StatusInternalServerError)
	}
}
