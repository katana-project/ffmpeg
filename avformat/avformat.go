package avformat

/*
#cgo pkg-config: libavformat

#include <libavformat/avformat.h>
*/
import "C"
import (
	"github.com/katana-project/ffmpeg/avutil"
	"unsafe"
)

type InputFormat struct {
	c *C.AVInputFormat
}

func (ifo *InputFormat) Unwrap() unsafe.Pointer {
	if ifo == nil {
		return nil
	}
	return unsafe.Pointer(ifo.c)
}

func (ifo *InputFormat) UnwrapDest() unsafe.Pointer {
	if ifo == nil {
		return nil
	}
	return unsafe.Pointer(&ifo.c)
}

type OutputFormat struct {
	c *C.AVOutputFormat
}

func (of *OutputFormat) Unwrap() unsafe.Pointer {
	if of == nil {
		return nil
	}
	return unsafe.Pointer(of.c)
}

func (of *OutputFormat) UnwrapDest() unsafe.Pointer {
	if of == nil {
		return nil
	}
	return unsafe.Pointer(&of.c)
}

type FormatContext struct {
	c *C.AVFormatContext
}

func (fc *FormatContext) Unwrap() unsafe.Pointer {
	if fc == nil {
		return nil
	}
	return unsafe.Pointer(fc.c)
}

func (fc *FormatContext) UnwrapDest() unsafe.Pointer {
	if fc == nil {
		return nil
	}
	return unsafe.Pointer(&fc.c)
}

func (fc *FormatContext) OpenInput(url string, fmt *InputFormat, options *avutil.Dictionary) int {
	url0 := C.CString(url)
	defer C.free(unsafe.Pointer(url0))

	return int(C.avformat_open_input(&fc.c, url0, (*C.AVInputFormat)(fmt.Unwrap()), (**C.AVDictionary)(options.UnwrapDest())))
}

func (fc *FormatContext) FindStreamInfo(options *avutil.Dictionary) int {
	return int(C.avformat_find_stream_info(fc.c, (**C.AVDictionary)(options.UnwrapDest())))
}

func (fc *FormatContext) DumpFormat(index int, url string, isOutput bool) {
	url0 := C.CString(url)
	defer C.free(unsafe.Pointer(url0))

	isOutput0 := 0
	if isOutput { // boolean to int
		isOutput0 = 1
	}

	C.av_dump_format(fc.c, C.int(index), url0, C.int(isOutput0))
}

func (fc *FormatContext) AllocOutputContext2(oformat *OutputFormat, formatName, filename string) int {
	var (
		formatName0 = C.CString(formatName)
		filename0   = C.CString(filename)
	)
	defer C.free(unsafe.Pointer(formatName0))
	defer C.free(unsafe.Pointer(filename0))

	return int(C.avformat_alloc_output_context2(&fc.c, (*C.AVOutputFormat)(oformat.Unwrap()), formatName0, filename0))
}
