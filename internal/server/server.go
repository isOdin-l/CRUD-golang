package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

func New() Server {
	return Server{httpServer: &http.Server{}}
}

func (s *Server) Run(port string, router chi.Router) error {
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        router,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) GracefulShutdown(ctx context.Context) {
	quitChannel := make(chan os.Signal, 1)

	signal.Notify(quitChannel, syscall.SIGTERM, syscall.SIGINT)
	<-quitChannel

	logrus.Print("Server is shutting down")

	err := s.httpServer.Shutdown(ctx)

	if err != http.ErrServerClosed && err != nil {
		logrus.Errorf("error on server shutting down: %s", err.Error())
	} else {
		logrus.Info("Server grecefully stoped")
	}
}
