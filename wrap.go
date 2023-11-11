package ffmpeg

import "unsafe"

type ValueWrapper interface {
	Unwrap() unsafe.Pointer
}

type PointerWrapper interface {
	ValueWrapper

	Null() bool
	UnwrapDest() unsafe.Pointer
}
