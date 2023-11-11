package avcodec

/*
#cgo pkg-config: libavcodec

#include <libavcodec/codec.h>
*/
import "C"
import "unsafe"

type Codec struct {
	c *C.AVCodec
}

func NewCodec(ptr unsafe.Pointer) *Codec {
	return &Codec{c: (*C.AVCodec)(ptr)}
}

func (c *Codec) Null() bool {
	return c.c == nil
}

func (c *Codec) Unwrap() unsafe.Pointer {
	if c == nil {
		return nil
	}
	return unsafe.Pointer(c.c)
}

func (c *Codec) UnwrapDest() unsafe.Pointer {
	if c == nil {
		return nil
	}
	return unsafe.Pointer(&c.c)
}
