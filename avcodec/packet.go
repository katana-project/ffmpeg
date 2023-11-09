package avcodec

/*
#cgo pkg-config: libavcodec

#include <libavcodec/packet.h>
*/
import "C"
import "unsafe"

type Packet struct {
	c *C.AVPacket
}

func NewPacket(ptr unsafe.Pointer) *Packet {
	return &Packet{c: (*C.AVPacket)(ptr)}
}

func (p *Packet) Alloc() {
	p.c = C.av_packet_alloc()
	if p.c == nil {
		panic("av_packet_alloc failed")
	}
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

func (p *Packet) Free() {
	C.av_packet_free(&p.c)
}
