package model

import (
	"time"

	"github.com/google/uuid"
)

// User defines domain model json and db representation
type User struct {
	UID       uuid.UUID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Cpf       string    `db:"cpf" json:"cpf"`
	BirthDate time.Time `db:"birthdate" json:"birthdate"`
}
