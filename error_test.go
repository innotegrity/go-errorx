package errorx_test

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
	"time"

	"go.innotegrity.dev/errorx"
)

// TODO: implement testing and benchmarks

type GenericError struct{ errorx.BaseError }

func NewGenericError(err error) GenericError {
	return GenericError{
		BaseError: errorx.BaseError{
			ErrCode: 128,
			Err:     err,
		},
	}
}

func (e GenericError) Code() int {
	if e.ErrCode == 0 {
		return 128
	}
	return e.ErrCode
}

func (e GenericError) Error() string {
	buf := bytes.NewBuffer(nil)
	if e.Err == nil {
		fmt.Fprintf(buf, "an unknown error occurred (code=%d)", e.Code())
	} else {
		fmt.Fprintf(buf, "%s (code=%d)", e.Err.Error(), e.Code())
	}
	if len(e.ErrAttrs) > 0 {
		buf.WriteString(" [")
		for k, v := range e.ErrAttrs {
			fmt.Fprintf(buf, " %s=%v", k, v)
		}
		buf.WriteString(" ]")
	}
	for _, n := range e.NestedErrors() {
		buf.WriteString("\n   ")
		buf.WriteString(n.Error())
	}
	return buf.String()
}

func TestBaseError1(t *testing.T) {
	err := errors.New("this is an error")
	e := NewGenericError(err)
	t.Logf("error: %s\n", e.Error())
	e2 := NewGenericError(err)
	e2.ErrCode = 1234
	t.Logf("error: %s\n", e2.Error())
	e3 := GenericError{
		BaseError: errorx.BaseError{
			Err: err,
			ErrAttrs: map[string]any{
				"key1": "value1",
				"key2": 2334,
				"key3": time.Now().UTC(),
			},
			NestedErrs: []errorx.Error{e, e2},
		},
	}
	t.Logf("error: %s\n", e3.Error())
}
