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

func (cc *CodecContext) Framerate() *avutil.Rational {
	return avutil.NewRational(unsafe.Pointer(&cc.c.framerate))
}

func (cc *CodecContext) SetFramerate(framerate *avutil.Rational) {
	cc.c.framerate = *(*C.AVRational)(framerate.Unwrap())
}

func (cc *CodecContext) PixFmt() avutil.PixelFormat {
	return avutil.PixelFormat(cc.c.pix_fmt)
}

func (cc *CodecContext) SetPixFmt(pixFmt avutil.PixelFormat) {
	cc.c.pix_fmt = int32(pixFmt)
}

func (cc *CodecContext) Width() int {
	return int(cc.c.width)
}

func (cc *CodecContext) SetWidth(width int) {
	cc.c.width = C.int(width)
}

func (cc *CodecContext) Height() int {
	return int(cc.c.height)
}

func (cc *CodecContext) SetHeight(height int) {
	cc.c.height = C.int(height)
}

func (cc *CodecContext) DecodeSubtitle2(sub *Subtitle, avpkt *Packet) (int, int) {
	var gotSubPtr C.int
	return int(C.avcodec_decode_subtitle2(cc.c, sub.c, &gotSubPtr, avpkt.c)), int(gotSubPtr)
}

func (cc *CodecContext) EncodeSubtitle(pkt *Packet, sub *Subtitle) int { // deviation from libavcodec's api - use a packet's buffer
	return int(C.avcodec_encode_subtitle(cc.c, pkt.c.data, pkt.c.size, sub.c))
}

func (cc *CodecContext) SendFrame(frame *avutil.Frame) int {
	return int(C.avcodec_send_frame(cc.c, (*C.AVFrame)(frame.Unwrap())))
}

func (cc *CodecContext) ReceiveFrame(frame *avutil.Frame) int {
	return int(C.avcodec_receive_frame(cc.c, (*C.AVFrame)(frame.Unwrap())))
}

func (cc *CodecContext) SendPacket(avpkt *Packet) int {
	return int(C.avcodec_send_packet(cc.c, avpkt.c))
}

func (cc *CodecContext) ReceivePacket(avpkt *Packet) int {
	return int(C.avcodec_receive_packet(cc.c, avpkt.c))
}

func (cc *CodecContext) Close() int {
	return int(C.avcodec_close(cc.c))
}

func (cc *CodecContext) FreeContext() {
	C.avcodec_free_context(&cc.c)
}
