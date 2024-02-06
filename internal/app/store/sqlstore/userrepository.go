package sqlstore

import (
	"database/sql"

	model "github.com/Andrew-Savin-msk/http-rest-api/internal/app/model/user"
	"github.com/Andrew-Savin-msk/http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
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

	// Insert into table
	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING user_id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID)
}

func (r *UserRepository) Find(id int) (*model.User, error) {
	u := &model.User{}

	err := r.store.db.QueryRow(
		"SELECT user_id, email, encrypted_password FROM users WHERE user_id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}

	err := r.store.db.QueryRow(
		"SELECT user_id, email, encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}
