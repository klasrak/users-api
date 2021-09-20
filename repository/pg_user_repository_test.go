package repository

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	model "github.com/klasrak/users-api/models"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
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
			uid, _ := uuid.NewRandom()
			u := &model.User{
				UID:       uid,
				Name:      faker.Name(),
				Email:     faker.Email(),
				Cpf:       "313.716.772-80",
				BirthDate: time.Date(1990, 1, 1, 1, 1, 1, 1, time.UTC),
			}
			db, mock := NewMock()

			sqlxDB := sqlx.NewDb(db, "sqlmock")

			defer sqlxDB.Close()

			query := `SELECT \* FROM users u;`

			userRepository := &UserRepository{DB: sqlxDB}

			rows := sqlmock.NewRows([]string{"id", "name", "email", "cpf", "birthdate"}).AddRow(u.UID, u.Name, u.Email, u.Cpf, u.BirthDate)

			mock.ExpectQuery(query).WillReturnRows(rows)

			ctx := context.Background()

			users, err := userRepository.GetAll(ctx, "")

			assert.NotNil(t, users)
			assert.NoError(t, err)
			assert.IsType(t, []model.User{}, users)
			assert.Len(t, users, 1)
		})
	})
}
