package errorx

import (
	"fmt"
	"time"
)

// Error represents an extension to the standard error interface by adding the ability to include an error code,
// nested errors, and any attributes associated with the error.
type Error interface {
	// Attrs should return any additional attributes which may be used when logging the error.
	Attrs() map[string]any

	// Code should return the corresponding error code.
	Code() int

	// Error should return the string version of the error.
	Error() string

	// InternalError should return the internal standard error object if there is one.
	InternalError() error

	// NestedErrors should return the list of nested errors if there are any.
	NestedErrors() []Error
}

// BaseError is a base error object that implements the Error interface and can be used to more easily compose more
// type-specific errors.
type BaseError struct {
	// ErrAttrs is a map of attributes to associate with the error itself.
	ErrAttrs map[string]any

	// ErrCode is the code to use for the error.
	ErrCode int

	// Err is the actual standard error that occurred.
	Err error

	// NestedErrors holds an array of errors that were generated by a call to another function.
	NestedErrs []Error
}

// Attr returns the value of the attribute with the given key if it exists.
func (b BaseError) Attr(key string) (any, error) {
	v, ok := b.ErrAttrs[key]
	if !ok {
		return "", fmt.Errorf("'%s': attribute not found", key)
	}
	return v, nil
}

// AttrDuration returns the value of the attribute with the given key (if it exists) as a time duration.
func (b BaseError) AttrDuration(key string) (time.Duration, error) {
	v, err := b.Attr(key)
	if err != nil {
		return time.Duration(0), err
	}
	t, ok := v.(time.Duration)
	if !ok {
		return time.Duration(0), fmt.Errorf("'%s': cannot convert attribute value to time.Duration", key)
	}
	return t, nil
}

// AttrInt returns the value of the attribute with the given key (if it exists) as an integer if it exists.
func (b BaseError) AttrInt(key string) (int, error) {
	v, err := b.Attr(key)
	if err != nil {
		return 0, err
	}
	i, ok := v.(int)
	if !ok {
		return 0, fmt.Errorf("'%s': cannot convert attribute value to an int", key)
	}
	return i, nil
}

// AttrInt64 returns the value of the attribute with the given key (if it exists) as a 64-bit integer.
func (b BaseError) AttrInt64(key string) (int64, error) {
	v, err := b.Attr(key)
	if err != nil {
		return 0, err
	}
	i, ok := v.(int64)
	if !ok {
		return 0, fmt.Errorf("'%s': cannot convert attribute value to an int64", key)
	}
	return i, nil
}

// Attrs returns all of the attributes associated with the error.
func (b BaseError) Attrs() map[string]any {
	if b.ErrAttrs == nil {
		return map[string]any{}
	}
	return b.ErrAttrs
}

// AttrString returns the value of the attribute with the given key (if it exists) as a string.
func (b BaseError) AttrString(key string) (string, error) {
	v, err := b.Attr(key)
	if err != nil {
		return "", err
	}
	str, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("'%s': cannot convert attribute value to a string", key)
	}
	return str, nil
}

// AttrTime returns the value of the attribute with the given key (if it exists) as a time object.
func (b BaseError) AttrTime(key string) (time.Time, error) {
	v, err := b.Attr(key)
	if err != nil {
		return time.Time{}, err
	}
	t, ok := v.(time.Time)
	if !ok {
		return time.Time{}, fmt.Errorf("'%s': cannot convert attribute value to time.Time", key)
	}
	return t, nil
}

// AttrUint returns the value of the attribute with the given key (if it exists) as an unsigned integer.
func (b BaseError) AttrUint(key string) (uint, error) {
	v, err := b.Attr(key)
	if err != nil {
		return 0, err
	}
	i, ok := v.(uint)
	if !ok {
		return 0, fmt.Errorf("'%s': cannot convert attribute value to an uint", key)
	}
	return i, nil
}

// AttrUint64 returns the value of the attribute with the given key (if it exists) as an unsigned 64-bit integer.
func (b BaseError) AttrUint64(key string) (uint64, error) {
	v, err := b.Attr(key)
	if err != nil {
		return 0, err
	}
	i, ok := v.(uint64)
	if !ok {
		return 0, fmt.Errorf("'%s': cannot convert attribute value to an uint64", key)
	}
	return i, nil
}

// Code returns the error code.
func (b BaseError) Code() int {
	return b.ErrCode
}

// Error returns a string representation of the error.
func (b BaseError) Error() string {
	if b.Err == nil {
		return "an unknown error occurred"
	}
	return fmt.Sprintf("error: %s", b.Err.Error())
}

// InternalError returns the standard error associated with the object.
func (b BaseError) InternalError() error {
	return b.Err
}

// NestedErrors returns the list of errors that were generated by a call to another function.
func (b BaseError) NestedErrors() []Error {
	if b.NestedErrs == nil {
		return []Error{}
	}
	return b.NestedErrs
}