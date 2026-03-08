package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Maxagena/sprint13/pkg/api"
)

type Server struct {
	Logger     *log.Logger
	HTTPServer *http.Server
}

func NewServer(logger *log.Logger) *Server {

	api.Init()

	webDir := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	httpServer := &http.Server{
		Addr:         ":7540",
		Handler:      nil,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     logger,
		HTTPServer: httpServer,
	}
}
