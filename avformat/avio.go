package avformat

/*
#cgo pkg-config: libavformat

#include <stdlib.h>
#include <libavformat/avio.h>
*/
import "C"
import (
	"unsafe"
)

const (
	IOFlagRead      = 1
	IOFlagWrite     = 2
	IOFlagReadWrite = IOFlagRead | IOFlagWrite
)

type IOContext struct {
	c *C.AVIOContext
}

func NewIOContext(ptr unsafe.Pointer) *IOContext {
	return &IOContext{c: (*C.AVIOContext)(ptr)}
}

func (ioc *IOContext) Null() bool {
	return ioc.c == nil
}

func (ioc *IOContext) Unwrap() unsafe.Pointer {
	if ioc == nil {
		return nil
	}
	return unsafe.Pointer(ioc.c)
}

func (ioc *IOContext) UnwrapDest() unsafe.Pointer {
	if ioc == nil {
		return nil
	}
	return unsafe.Pointer(&ioc.c)
}

func (ioc *IOContext) Open(url string, flags int) int {
	url0 := C.CString(url)
	defer C.free(unsafe.Pointer(url0))

	return int(C.avio_open(&ioc.c, url0, C.int(flags)))
}

func (ioc *IOContext) CloseP() int {
	return int(C.avio_closep(&ioc.c))
}
