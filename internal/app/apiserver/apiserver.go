package apiserver

import "github.com/sirupsen/logrus"

type APIServer struct {
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
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
