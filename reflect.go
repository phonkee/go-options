package options

import "reflect"

// isPointer checks if the type is a pointer.
func isPointer[T any]() bool {
	typ := reflect.TypeFor[T]()
	if typ == nil {
		return false
	}
	return typ.Kind() == reflect.Ptr
}
