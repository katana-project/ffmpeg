package avcodec

/*
#cgo pkg-config: libavcodec

#include <libavcodec/avcodec.h>
*/
import "C"
import (
	"github.com/katana-project/ffmpeg/avutil"
	"unsafe"
)

type CodecContext struct {
	c *C.AVCodecContext
}

func (cc *CodecContext) Null() bool {
	return cc.c == nil
}

func (cc *CodecContext) Unwrap() unsafe.Pointer {
	if cc == nil {
		return nil
	}
	return unsafe.Pointer(cc.c)
}

func (cc *CodecContext) UnwrapDest() unsafe.Pointer {
	if cc == nil {
		return nil
	}
	return unsafe.Pointer(&cc.c)
}

func (cc *CodecContext) AllocContext3(codec *Codec) {
	cc.c = C.avcodec_alloc_context3(codec.c)
}

func (cc *CodecContext) Open2(codec *Codec, options *avutil.Dictionary) int {
	return int(C.avcodec_open2(cc.c, codec.c, (**C.AVDictionary)(options.UnwrapDest())))
}

func (cc *CodecContext) ParametersFromContext(par *CodecParameters) int {
	return int(C.avcodec_parameters_from_context(par.c, cc.c))
}

func (cc *CodecContext) ParametersToContext(par *CodecParameters) int {
	return int(C.avcodec_parameters_to_context(cc.c, par.c))
}

func (cc *CodecContext) FreeContext() {
	C.avcodec_free_context(&cc.c)
}
