package options

import (
	"fmt"
)

// New creates a new Options instance.
func New[T any](value T, opts ...Option[T]) (_ *Options[T], err error) {
	result := &Options[T]{
		value: value,
	}

	// we do not allow pointer types
	if isPointer[T]() {
		return nil, fmt.Errorf("%w: options cannot be pointer", ErrImproperlyConfigured)
	}

	// Apply given options
	if err = result.Apply(opts...); err != nil {
		return nil, err
	}

	return result, nil
}

// Options is a struct that holds inner options as value
type Options[T any] struct {
	value T
}

// Apply applies the options to the given value.
func (o *Options[T]) Apply(opts ...Option[T]) error {
	for _, opt := range opts {
		if err := opt(&o.value); err != nil {
			return err
		}
	}
	return nil
}

// Get returns the value of the options.
func (o *Options[T]) Get() T {
	return o.value
}

// Validate validates the options.
func (o *Options[T]) Validate() error {
	if any(o.value) == nil {
		return nil
	}

	// check if we have value and if it implements Validator
	if any(o.value) != nil {
		if i, ok := any(&o.value).(Validator); ok {
			return i.Validate()
		}
	}

	return nil
}
