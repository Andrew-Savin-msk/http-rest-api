package apiserver

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Andrew-Savin-msk/http-rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// added router on 17:48
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	const op = "internal.app.apiserver.Start"
	err := s.configureLogger()
	if err != nil {
		return fmt.Errorf("%s; error with starting server in: %w", op, err)
	}

	s.configureRouter()

	err = s.configureStore()

	if err != nil {
		return fmt.Errorf("%s; error with starting server in: %w", op, err)
	}

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	const op = "internal.app.apiserver.configureLogger"
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return fmt.Errorf("%s, with setting configure level, ended with error: %w", op, err)
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) configureStore() error {
	const op = "internal.app.apiserver.configureStore"
	st := store.New(s.config.Store)
	err := st.Open()
	if err != nil {
		return fmt.Errorf("%s, with opening store, ended with error: %w", op, err)
	}

	s.store = st

	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
