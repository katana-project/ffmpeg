package avutil

/*
#cgo pkg-config: libavutil

#include <libavutil/frame.h>
*/
import "C"
import "unsafe"

type Frame struct {
	c *C.AVFrame
}

func NewFrame(ptr unsafe.Pointer) *Frame {
	return &Frame{c: (*C.AVFrame)(ptr)}
}

func (f *Frame) Null() bool {
	return f.c == nil
}

func (f *Frame) Unwrap() unsafe.Pointer {
	if f == nil {
		return nil
	}
	return unsafe.Pointer(f.c)
}

func (f *Frame) UnwrapDest() unsafe.Pointer {
	if f == nil {
		return nil
	}
	return unsafe.Pointer(&f.c)
}

func (f *Frame) Alloc() {
	f.c = C.av_frame_alloc()
}

func (f *Frame) Free() {
	C.av_frame_free(&f.c)
}

func (f *Frame) Unref() {
	C.av_frame_unref(f.c)
}

func (f *Frame) PictType() PictureType {
	return PictureType(f.c.pict_type)
}

func (f *Frame) SetPictType(pictType PictureType) {
	f.c.pict_type = uint32(pictType)
}
