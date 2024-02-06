package teststore

import (
	model "github.com/Andrew-Savin-msk/http-rest-api/internal/app/model/user"
	"github.com/Andrew-Savin-msk/http-rest-api/internal/app/store"
)

type Store struct {
	userRepository *UserRepository
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
