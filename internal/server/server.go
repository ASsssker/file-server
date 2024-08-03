package server

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	*http.Server
}

func NewServer(host, path string) *Server {
	handler := NewHandler(path)

	return &Server{
		&http.Server{
			Addr:         host,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
			Handler:      handler,
		}}
}

func (s *Server) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("Ошибка при запуске сервера: %s\n", err.Error())
		}
	}()

	log.Printf("Сервер слушает хост %s", s.Addr)
	<-ctx.Done()
	log.Println("Отключение сервера")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Stop(shutdownCtx); err != nil {
		log.Fatalf("Ошибка при отключении сервера: %s\n", err.Error())
	}
}

func (s Server) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}
