package store

import model "github.com/Andrew-Savin-msk/http-rest-api/internal/app/model/user"

type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}
