package avcodec

/*
#cgo pkg-config: libavcodec

#include <libavcodec/packet.h>
*/
import "C"
import (
	"github.com/katana-project/ffmpeg/avutil"
	"unsafe"
)

type Packet struct {
	c *C.AVPacket
}

func NewPacket(ptr unsafe.Pointer) *Packet {
	return &Packet{c: (*C.AVPacket)(ptr)}
}

func (p *Packet) Null() bool {
	return p.c == nil
}

func (p *Packet) Unwrap() unsafe.Pointer {
	if p == nil {
		return nil
	}
	return unsafe.Pointer(p.c)
}

func (p *Packet) UnwrapDest() unsafe.Pointer {
	if p == nil {
		return nil
	}
	return unsafe.Pointer(&p.c)
}

func (p *Packet) Alloc() {
	p.c = C.av_packet_alloc()
}

func (p *Packet) New(size int) int {
	return int(C.av_new_packet(p.c, C.int(size)))
}

func (p *Packet) Shrink(size int) {
	C.av_shrink_packet(p.c, C.int(size))
}

func (p *Packet) Grow(growBy int) int {
	return int(C.av_grow_packet(p.c, C.int(growBy)))
}

func (p *Packet) StreamIndex() int {
	return int(p.c.stream_index)
}

func (p *Packet) SetStreamIndex(index int) {
	p.c.stream_index = C.int(index)
}

func (p *Packet) Unref() {
	C.av_packet_unref(p.c)
}

func (p *Packet) RescaleTs(tbSrc, tbDst *avutil.Rational) {
	C.av_packet_rescale_ts(p.c, *(*C.AVRational)(tbSrc.Unwrap()), *(*C.AVRational)(tbDst.Unwrap()))
}

func (p *Packet) Pos() int64 {
	return int64(p.c.pos)
}

func (p *Packet) SetPos(pos int64) {
	p.c.pos = C.int64_t(pos)
}

func (p *Packet) Pts() int64 {
	return int64(p.c.pts)
}

func (p *Packet) Dts() int64 {
	return int64(p.c.dts)
}

func (p *Packet) Duration() int64 {
	return int64(p.c.duration)
}

func (p *Packet) Free() {
	C.av_packet_free(&p.c)
}
