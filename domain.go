package options

import "errors"

// Option is a function that is functional option for configuring a type.
type Option[T any] func(*T) error

// Validator is an interface that defines a method for validating options.
type Validator interface {
	Validate() error
}

var (
	ErrImproperlyConfigured = errors.New("improperly configured")
)
