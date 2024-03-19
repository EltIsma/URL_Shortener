package server

import (
	"context"
	"net/http"
	"time"
	"urlShortener/urlShortener/config"
)

const (
	defaultReadTimeout  = 30 * time.Second
	defaultWriteTimeout = 30 * time.Second
)

type Server struct {
	server *http.Server
}

func (s *Server) New(handler http.Handler, config *config.Config) error {
	readTimeout := config.ReadTimeout
	if readTimeout < 1 {
		readTimeout = defaultReadTimeout
	}

	writeTimeout := config.WriteTimeout
	if writeTimeout < 1 {
		writeTimeout = defaultWriteTimeout
	}
	s.server = &http.Server{
		Addr:         config.AppHost + ":" + config.AppPort,
		Handler:      handler,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
