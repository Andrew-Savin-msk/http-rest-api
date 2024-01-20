package store

import (
	"github.com/Andrew-Savin-msk/http-rest-api/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	// Check if data valid for insert
	err := u.Validate()
	if err != nil {
		return nil, err
	}

	// Making all fields ready before usage
	err = u.BeforCreate()
	if err != nil {
		return nil, err
	}

	// Insert into table
	err = r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)

	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	)

	if err != nil {
		return nil, err
	}
	return u, nil
}
