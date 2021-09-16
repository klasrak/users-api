package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBrazilianCPFValidWithValidCPF(t *testing.T) {
	assert := assert.New(t)

	assert.True(IsBrazilianCPFValid("313.716.772-80"), "masked valid cpf, should be true")
	assert.True(IsBrazilianCPFValid("31371677280"), "unmasked valid cpf, should be true")
}

func TestIsBrazilianCPFValidWithInvalidCPF(t *testing.T) {
	assert := assert.New(t)

	assert.False(IsBrazilianCPFValid("111.111.111-11"), "masked invalid cpf, should be false")
	assert.False(IsBrazilianCPFValid("11111111111"), "unmasked invalid cpf, should be false")

	assert.False(IsBrazilianCPFValid("123.456.789-10"), "masked invalid cpf, should be false")
	assert.False(IsBrazilianCPFValid("12345678910"), "unmasked invalid cpf, should be false")

	assert.False(IsBrazilianCPFValid("918.577.800-18"), "masked invalid cpf, should be false")
	assert.False(IsBrazilianCPFValid("91857780018"), "unmasked invalid cpf, should be false")
}
