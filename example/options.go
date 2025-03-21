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

package main

import (
	"fmt"
	"github.com/phonkee/options"
)

// Option is an alias for options.Option
type Option = options.Option[opts]

// WithHello adds a hello option to the options and also applies another option
func WithHello(hello string) Option {
	return func(o *opts) error {
		if hello == "" {
			return fmt.Errorf("%w: hello cannot be empty", options.ErrImproperlyConfigured)
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
