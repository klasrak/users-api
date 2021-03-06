package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeBetween(t *testing.T) {
	assert := assert.New(t)

	past := time.Date(1993, 7, 7, 0, 0, 0, 0, time.UTC)
	present := time.Date(2021, 7, 7, 0, 0, 0, 0, time.UTC)

	year, month, day, hour, min, sec := TimeBetween(past, present)

	assert.Equal(28, year)
	assert.Equal(0, month)
	assert.Equal(0, day)
	assert.Equal(0, hour)
	assert.Equal(0, min)
	assert.Equal(0, sec)

}

func TestIsUnderage(t *testing.T) {
	t.Run("True", func(t *testing.T) {
		birthdate := time.Date(2006, 1, 1, 1, 1, 1, 1, time.UTC)

		isUnderage := IsUnderage(birthdate)

		assert.True(t, isUnderage)
	})

	t.Run("False", func(t *testing.T) {
		birthdate := time.Date(1993, 1, 1, 1, 1, 1, 1, time.UTC)

		isUnderage := IsUnderage(birthdate)

		assert.False(t, isUnderage)
	})
}
