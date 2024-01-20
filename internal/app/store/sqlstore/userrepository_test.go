package sqlstore_test

import (
	"testing"

	"github.com/Andrew-Savin-msk/http-rest-api/internal/app/model"
	"github.com/Andrew-Savin-msk/http-rest-api/internal/app/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("users")

	s := store.NewStore(db)
	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("users")

	s := store.NewStore(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = "user@example.org"

	_, err = s.User().Create(u)
	assert.NoError(t, err)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
