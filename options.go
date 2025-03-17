/*
 * MIT License
 *
 * Copyright (c) 2025 Peter Vrba
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

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
