package repository

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestUserRepository(t *testing.T) {
	t.Run("GetAll", func(t *testing.T) {
		t.Run("Success without name filter", func(t *testing.T) {
			// db, mock := NewMock()

			// sqlxDB := sqlx.NewDb(db, "sqlmock")

			// defer sqlxDB.Close()

			// userRepository := &UserRepository{DB: sqlxDB}

			// var users []model.User

			// err := faker.FakeData(&users)

			// assert.NoError(t, err)
		})
	})
}
