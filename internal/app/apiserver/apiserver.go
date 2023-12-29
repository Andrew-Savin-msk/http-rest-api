package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// added router on 17:48
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

func (s *APIServer) Start() error {
	err := s.configureLogger()
	if err != nil {
		return err
	}

	s.logger.Info("starting api server")

	return nil
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}
