package apiserver

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (server *APIServer) Start() error {
	if err := server.ConfigureLogger(); err != nil {
		return err
	}

	server.ConfigureRouter()
	server.logger.Info(fmt.Sprintf("Server started at: %s", server.config.Address))

	return http.ListenAndServe(server.config.Address, server.router)
}

func (server *APIServer) ConfigureLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)

	if err != nil {
		return err
	}

	server.logger.SetLevel(level)
	return nil
}

func (server *APIServer) ConfigureRouter() {
	server.router.HandleFunc("/hello", server.HandleHello())
}

func (server *APIServer) HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello!")
	}
}
