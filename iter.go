package ffmpeg

import "unsafe"

// wrapper for void** types

type IterationState struct {
	c unsafe.Pointer
}

func NewIterationState(ptr unsafe.Pointer) *IterationState {
	return &IterationState{c: ptr}
}

func (is *IterationState) Null() bool {
	return is.c == nil
}

func (is *IterationState) Unwrap() unsafe.Pointer {
	if is == nil {
		return nil
	}
	return is.c
}

func (is *IterationState) UnwrapDest() unsafe.Pointer {
	if is == nil {
		return nil
	}
	return unsafe.Pointer(&is.c)
}
