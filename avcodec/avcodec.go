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

type Subtitle struct {
	c *C.AVSubtitle
}

func NewSubtitle(ptr unsafe.Pointer) *Subtitle {
	return &Subtitle{c: (*C.AVSubtitle)(ptr)}
}

func (s *Subtitle) Null() bool {
	return s.c == nil
}

func (s *Subtitle) Unwrap() unsafe.Pointer {
	if s == nil {
		return nil
	}
	return unsafe.Pointer(s.c)
}

func (s *Subtitle) UnwrapDest() unsafe.Pointer {
	if s == nil {
		return nil
	}
	return unsafe.Pointer(&s.c)
}

type CodecContext struct {
	c *C.AVCodecContext
}

func NewCodecContext(ptr unsafe.Pointer) *CodecContext {
	return &CodecContext{c: (*C.AVCodecContext)(ptr)}
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

func (cc *CodecContext) TimeBase() *avutil.Rational {
	return avutil.NewRational(unsafe.Pointer(&cc.c.time_base))
}

func (cc *CodecContext) SetTimeBase(timeBase *avutil.Rational) {
	cc.c.time_base = *(*C.AVRational)(timeBase.Unwrap())
}

func (cc *CodecContext) DecodeSubtitle2(sub *Subtitle, avpkt *Packet) (int, int) {
	var gotSubPtr C.int
	return int(C.avcodec_decode_subtitle2(cc.c, sub.c, &gotSubPtr, avpkt.c)), int(gotSubPtr)
}

func (cc *CodecContext) EncodeSubtitle(pkt *Packet, sub *Subtitle) int { // deviation from libavcodec's api - use a packet's buffer
	return int(C.avcodec_encode_subtitle(cc.c, pkt.c.data, pkt.c.size, sub.c))
}

func (cc *CodecContext) Close() int {
	return int(C.avcodec_close(cc.c))
}

func (cc *CodecContext) FreeContext() {
	C.avcodec_free_context(&cc.c)
}
