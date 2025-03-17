package main

import (
	"github.com/phonkee/options"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("test valid", func(t *testing.T) {
		o, err := options.New(opts{
			Hello: "world",
		},
			WithHello("hello"),
		)
		assert.NoError(t, err)
		assert.Equal(t, "hello", o.Get().Hello)
		assert.Equal(t, "world", o.Get().World)

	})

	t.Run("test invalid", func(t *testing.T) {
		o, err := options.New(opts{
			Hello: "hello",
		},
			WithHello(""),
		)
		assert.Nil(t, o)
		assert.ErrorIs(t, err, ErrInvalidHello)
		assert.ErrorIs(t, err, options.ErrImproperlyConfigured)

	})
}
