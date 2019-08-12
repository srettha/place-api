package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadEnv(t *testing.T) {
	err := LoadEnv()
	assert.Nil(t, err)
	assert.Equal(t, "place-api", Values.Name)
	assert.Equal(t, 3000, Values.Port)
	assert.Equal(t, "postgres", Values.Database.Dialect)
	assert.Equal(t, "postgres", Values.Database.Name)
}

