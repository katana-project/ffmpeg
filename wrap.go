package ffmpeg

import "unsafe"

type UnsafeWrapper interface {
	Null() bool
	Unwrap() unsafe.Pointer
	UnwrapDest() unsafe.Pointer
}
