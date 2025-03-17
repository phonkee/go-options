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
