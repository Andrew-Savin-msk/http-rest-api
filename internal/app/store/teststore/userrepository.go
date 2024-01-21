package teststore

import (
	"github.com/Andrew-Savin-msk/http-rest-api/internal/app/model"
	"github.com/Andrew-Savin-msk/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	// Check if data valid for insert
	err := u.Validate()
	if err != nil {
		return err
	}

	// Making all fields ready before usage
	err = u.BeforCreate()
	if err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u, ok := r.users[email]

	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}
