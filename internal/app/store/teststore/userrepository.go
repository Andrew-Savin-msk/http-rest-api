package teststore

import (
	model "github.com/Andrew-Savin-msk/http-rest-api/internal/app/model/user"
	"github.com/Andrew-Savin-msk/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
	users map[int]*model.User
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

	u.ID = len(r.users) + 1
	r.users[u.ID] = u
	u.ID = len(r.users)
	return nil
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u, ok := r.users[id]

	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, store.ErrRecordNotFound
}
