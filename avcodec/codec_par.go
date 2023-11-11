package avcodec

/*
#cgo pkg-config: libavcodec

#include <libavcodec/codec_par.h>
*/
import "C"
import (
	"github.com/katana-project/ffmpeg/avutil"
	"unsafe"
)

type CodecParameters struct {
	c *C.AVCodecParameters
}

func NewCodecParameters(ptr unsafe.Pointer) *CodecParameters {
	return &CodecParameters{c: (*C.AVCodecParameters)(ptr)}
}

func (cp *CodecParameters) Null() bool {
	return cp.c == nil
}

func (cp *CodecParameters) Unwrap() unsafe.Pointer {
	if cp == nil {
		return nil
	}
	return unsafe.Pointer(cp.c)
}

func (cp *CodecParameters) UnwrapDest() unsafe.Pointer {
	if cp == nil {
		return nil
	}
	return unsafe.Pointer(&cp.c)
}

func (cp *CodecParameters) CodecType() avutil.MediaType {
	return avutil.MediaType(cp.c.codec_type)
}

func (cp *CodecParameters) Copy(src *CodecParameters) int {
	return int(C.avcodec_parameters_copy(cp.c, src.c))
}

func (cp *CodecParameters) SetCodecTag(codecTag uint32) {
	cp.c.codec_tag = C.uint32_t(codecTag)
}
