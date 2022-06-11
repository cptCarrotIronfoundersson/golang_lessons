package internalhttp

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/configs/config"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
)

type Server struct {
	Host string
	Port string
}

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
}

type Application interface { // TODO
}

func myHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
}

func NewServer(logger *logger.Logger, config *config.Config, app Application) *Server {
	logger.Info(fmt.Sprintf("Server startded:  Host %v, Port %v", config.Server.Host, config.Server.Port))
	fmt.Println(config)
	http.Handle("/", loggingMiddleware(logger, http.HandlerFunc(myHandler)))
	http.Handle("/hello-world", loggingMiddleware(logger, http.HandlerFunc(myHandler)))

	return &Server{
		Host: config.Server.Host,
		Port: config.Server.Port,
	}
}

func (s *Server) Start(ctx context.Context) error {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", s.Host, s.Port), nil)
	fmt.Println(err)
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	os.Exit(1)
	return nil
}

// TODO
