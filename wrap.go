package ffmpeg

import "unsafe"

type UnsafeWrapper interface {
	Unwrap() unsafe.Pointer
	UnwrapDest() unsafe.Pointer
}
