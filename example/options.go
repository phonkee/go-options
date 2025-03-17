package main

import (
	"fmt"
	"github.com/phonkee/options"
)

var (
	ErrInvalidHello = fmt.Errorf("%w: hello cannot be empty", options.ErrImproperlyConfigured)
)

// Option is an alias for options.Option
type Option = options.Option[opts]

// WithHello adds a hello option to the options and also applies another option
func WithHello(hello string) Option {
	return func(o *opts) error {
		if hello == "" {
			return ErrInvalidHello
		}
		o.Hello = hello
		// add another options here
		for _, opt := range []Option{
			WithWorld("world"),
		} {
			if err := opt(o); err != nil {
				return err
			}
		}

		return nil
	}
}

// WithWorld adds a world option to the options
func WithWorld(world string) Option {
	return func(o *opts) error {
		o.World = world
		return nil
	}
}

// opts are actual options
type opts struct {
	Hello string
	World string
}
