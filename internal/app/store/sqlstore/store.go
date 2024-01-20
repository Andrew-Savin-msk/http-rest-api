package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
