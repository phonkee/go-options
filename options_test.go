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
	"github.com/stretchr/testify/assert"
	"testing"
)

// ptrTo is a helper function to create a pointer to a value of type T
func ptrTo[T any](v T) *T {
	return &v
}

func TestNew(t *testing.T) {
	t.Run("invalid pointer type", func(t *testing.T) {
		opts, err := New[*int](ptrTo(1))
		assert.Nil(t, opts)
		assert.ErrorIs(t, err, ErrImproperlyConfigured)
	})

	t.Run("valid type", func(t *testing.T) {
		opts, err := New(42)
		assert.NotNil(t, opts)
		assert.NoError(t, err)
		assert.Equal(t, 42, opts.Get())
	})
}
