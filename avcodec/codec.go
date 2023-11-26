package avcodec

/*
#cgo pkg-config: libavcodec

#include <libavcodec/codec.h>
*/
import "C"
import (
	"github.com/katana-project/ffmpeg"
	"github.com/katana-project/ffmpeg/avutil"
	"unsafe"
)

type Codec struct {
	c *C.AVCodec
}

func NewCodec(ptr unsafe.Pointer) *Codec {
	return &Codec{c: (*C.AVCodec)(ptr)}
}

func (c *Codec) Null() bool {
	return c.c == nil
}

func (c *Codec) Unwrap() unsafe.Pointer {
	if c == nil {
		return nil
	}
	return unsafe.Pointer(c.c)
}

func (c *Codec) UnwrapDest() unsafe.Pointer {
	if c == nil {
		return nil
	}
	return unsafe.Pointer(&c.c)
}

func (c *Codec) Name() string {
	return C.GoString(c.c.name)
}

func (c *Codec) LongName() string {
	return C.GoString(c.c.long_name)
}

func (c *Codec) Type() avutil.MediaType {
	return avutil.MediaType(c.c._type)
}

func (c *Codec) ID() CodecID {
	return CodecID(c.c.id)
}

func (c *Codec) IsEncoder() int {
	return int(C.av_codec_is_encoder(c.c))
}

func (c *Codec) IsDecoder() int {
	return int(C.av_codec_is_decoder(c.c))
}

func CodecIterate(opaque *ffmpeg.IterationState) *Codec {
	c := C.av_codec_iterate((*unsafe.Pointer)(opaque.UnwrapDest()))
	if c == nil {
		return nil
	}

	return &Codec{c: c}
}

func FindDecoder(id CodecID) *Codec {
	c := C.avcodec_find_decoder(uint32(id))
	if c == nil {
		return nil
	}

	return &Codec{c: c}
}

func FindDecoderByName(name string) *Codec {
	name0 := C.CString(name)
	defer C.free(unsafe.Pointer(name0))

	c := C.avcodec_find_decoder_by_name(name0)
	if c == nil {
		return nil
	}

	return &Codec{c: c}
}

func FindEncoder(id CodecID) *Codec {
	c := C.avcodec_find_encoder(uint32(id))
	if c == nil {
		return nil
	}

	return &Codec{c: c}
}

func FindEncoderByName(name string) *Codec {
	name0 := C.CString(name)
	defer C.free(unsafe.Pointer(name0))

	c := C.avcodec_find_encoder_by_name(name0)
	if c == nil {
		return nil
	}

	return &Codec{c: c}
}
