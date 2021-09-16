package utils

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIsBrazilianCPFValid(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	t.Run("Success with valid masked cpf", func(t *testing.T) {
		assert := assert.New(t)

		assert.True(IsBrazilianCPFValid("313.716.772-80"), "masked valid cpf, should be true")
		assert.True(IsBrazilianCPFValid("648.173.761-39"), "masked valid cpf, should be true")
		assert.True(IsBrazilianCPFValid("417.653.125-82"), "masked valid cpf, should be true")
		assert.True(IsBrazilianCPFValid("656.387.324-38"), "masked valid cpf, should be true")
	})

	t.Run("Success with valid unmasked cpf", func(t *testing.T) {
		assert := assert.New(t)

		assert.True(IsBrazilianCPFValid("31371677280"), "unmasked valid cpf, should be true")
		assert.True(IsBrazilianCPFValid("64817376139"), "unmasked valid cpf, should be true")
		assert.True(IsBrazilianCPFValid("41765312582"), "unmasked valid cpf, should be true")
		assert.True(IsBrazilianCPFValid("65638732438"), "unmasked valid cpf, should be true")
	})

	t.Run("Invalid with masked cpf", func(t *testing.T) {
		assert := assert.New(t)

		assert.False(IsBrazilianCPFValid("313.716.772-85"), "masked invalid cpf, should be false")
		assert.False(IsBrazilianCPFValid("648.173.761-41"), "masked invalid cpf, should be false")
		assert.False(IsBrazilianCPFValid("417.653.125-00"), "masked invalid cpf, should be false")
		assert.False(IsBrazilianCPFValid("656.387.324-99"), "masked invalid cpf, should be false")
	})

	t.Run("Invalid with unmasked cpf", func(t *testing.T) {
		assert := assert.New(t)

		assert.False(IsBrazilianCPFValid("31371677285"), "unmasked invalid cpf, should be false")
		assert.False(IsBrazilianCPFValid("64817376141"), "unmasked invalid cpf, should be false")
		assert.False(IsBrazilianCPFValid("41765312500"), "unmasked invalid cpf, should be false")
		assert.False(IsBrazilianCPFValid("65638732499"), "unmasked invalid cpf, should be false")
	})
}
