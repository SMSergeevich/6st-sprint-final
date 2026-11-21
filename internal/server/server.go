package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server — обёртка над http.Server с логгером.
type Server struct {
	logger *log.Logger
	server *http.Server
}

// New создаёт http-роутер, регистрирует хендлеры
// и возвращает экземпляр структуры Server.
func New(logger *log.Logger) *Server {
	// 1. Создаём http-роутер (стандартная библиотека).
	mux := http.NewServeMux()

	// 2. Регистрируем хендлеры.
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/upload", handlers.Upload)

	// 3. Создаём экземпляр http.Server с настройками из ТЗ.
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		server: srv,
	}
}

// Start запускает сервер.
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
