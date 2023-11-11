package avutil

/*
#cgo pkg-config: libavutil

#include <libavutil/rational.h>
*/
import "C"
import (
	"unsafe"
)

type Rational struct {
	c C.AVRational
}

func NewRational(ptr unsafe.Pointer) *Rational {
	return &Rational{c: *(*C.AVRational)(ptr)}
}

func (r *Rational) Unwrap() unsafe.Pointer {
	if r == nil {
		return nil
	}
	return unsafe.Pointer(&r.c)
}
