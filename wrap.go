package ffmpeg

import "unsafe"

type UnsafeWrap interface {
	Unwrap() unsafe.Pointer
	UnwrapDest() unsafe.Pointer
}
