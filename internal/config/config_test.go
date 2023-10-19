package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"hukum-onlen-go/internal/config"
)

func TestNewConfig(t *testing.T) {
	t.Run("can't create new instance due to invalid env", func(t *testing.T) {
		cfg, err := config.NewConfig("an invalid file")

		assert.NotNil(t, err)
		assert.Nil(t, cfg)
	})

	t.Run("can't create new instance due to fail parsing env", func(t *testing.T) {
		cfg, err := config.NewConfig("../../test/fixture/env.invalid")

		assert.NotNil(t, err)
		assert.Nil(t, cfg)
	})

	t.Run("successfully create a new instance of Config", func(t *testing.T) {
		cfg, err := config.NewConfig("../../test/fixture/env.valid")

		assert.Nil(t, err)
		assert.NotNil(t, cfg)
	})
}
