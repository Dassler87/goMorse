package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server структура сервера с логгером и HTTP‑сервером
type Server struct {
	Logger *log.Logger
	Server *http.Server
}

// New создаёт и настраивает HTTP‑сервер
func New(logger *log.Logger) *Server {
	// Создаём HTTP‑роутер
	router := http.NewServeMux()

	// Регистрируем хендлеры
	router.HandleFunc("/", handlers.ServeIndex)
	router.HandleFunc("/upload", handlers.UploadHandler)

	// Создаём экземпляр HTTP‑сервера с настройками
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger: logger,
		Server: srv,
	}
}
