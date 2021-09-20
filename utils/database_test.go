package utils

import (
	"database/sql"
	"testing"
	"time"

	"github.com/klasrak/users-api/rerrors"
	"github.com/stretchr/testify/assert"
)

func TestSanitizeUpdateParams(t *testing.T) {
	t.Run("Should parse empty string to sql.NullString", func(t *testing.T) {
		user := struct {
			Name  string
			Email string
		}{
			"",
			"",
		}

		sanitizedUser, err := SanitizeUpdateParams(user)

		assert.NoError(t, err)
		assert.IsType(t, new(map[string]interface{}), &sanitizedUser)
		assert.Equal(t, sql.NullString{}, sanitizedUser["Name"])
		assert.Equal(t, sql.NullString{}, sanitizedUser["Email"])
	})

	t.Run("Should parse time.Time{} to sql.NullString", func(t *testing.T) {
		user := struct {
			Birthdate time.Time
		}{
			time.Time{},
		}

		sanitizedUser, err := SanitizeUpdateParams(user)

		assert.NoError(t, err)
		assert.IsType(t, new(map[string]interface{}), &sanitizedUser)
		assert.Equal(t, sql.NullString{}, sanitizedUser["Birthdate"])
	})

	t.Run("Do nothing with valid string", func(t *testing.T) {
		user := struct {
			Name string
		}{
			"John Doe",
		}

		sanitizedUser, err := SanitizeUpdateParams(user)

		assert.NoError(t, err)
		assert.IsType(t, new(map[string]interface{}), &sanitizedUser)
		assert.Equal(t, user.Name, sanitizedUser["Name"])
	})

	t.Run("Parse valid time.Time", func(t *testing.T) {
		user := struct {
			Birthdate time.Time
		}{
			time.Now(),
		}

		sanitizedUser, err := SanitizeUpdateParams(user)

		assert.NoError(t, err)
		assert.IsType(t, new(map[string]interface{}), &sanitizedUser)
		assert.Equal(t, user.Birthdate.Format(time.RFC3339), sanitizedUser["Birthdate"].(time.Time).Format(time.RFC3339))

	})

	t.Run("Default", func(t *testing.T) {
		user := struct {
			Age int
		}{
			18,
		}

		sanitizedUser, err := SanitizeUpdateParams(user)

		assert.NoError(t, err)
		assert.IsType(t, new(map[string]interface{}), &sanitizedUser)

		assert.Equal(t, float64(user.Age), sanitizedUser["Age"])
		assert.IsType(t, float64(user.Age), sanitizedUser["Age"])
	})

	t.Run("Internal Server Error", func(t *testing.T) {
		user := "invalid type"

		internalServerError := rerrors.NewInternal()

		sanitizedUser, err := SanitizeUpdateParams(user)

		assert.Error(t, err)
		assert.Nil(t, sanitizedUser)
		assert.Equal(t, internalServerError, err)
	})
}
