package avformat

/*
#cgo pkg-config: libavformat

#include <libavformat/avformat.h>
*/
import "C"
import (
	"github.com/katana-project/ffmpeg"
	"github.com/katana-project/ffmpeg/avcodec"
	"github.com/katana-project/ffmpeg/avutil"
	"unsafe"
)

const (
	FmtNoFile       = 0x0001
	FmtNeedNumber   = 0x0002
	FmtExperimental = 0x0004
	FmtShowIds      = 0x0008
	FmtGlobalHeader = 0x0040
	FmtNoTimestamps = 0x0080
	FmtGenericIndex = 0x0100
	FmtTsDiscont    = 0x0200
	FmtVariableFps  = 0x0400
	FmtNoDimensions = 0x0800
	FmtNoStreams    = 0x1000
	FmtNoBinSearch  = 0x2000
	FmtNoGenSearch  = 0x4000
	FmtNoByteSeek   = 0x8000
	FmtTsNonStrict  = 0x20000
	FmtTsNegative   = 0x40000
	FmtSeekToPts    = 0x4000000
)

type InputFormat struct {
	c *C.AVInputFormat
}

func NewInputFormat(ptr unsafe.Pointer) *InputFormat {
	return &InputFormat{c: (*C.AVInputFormat)(ptr)}
}

func (ifo *InputFormat) Null() bool {
	return ifo.c == nil
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

func (ifo *InputFormat) Flags() int {
	return int(ifo.c.flags)
}

func (ifo *InputFormat) Name() string {
	return C.GoString(ifo.c.name)
}

func (ifo *InputFormat) LongName() string {
	return C.GoString(ifo.c.long_name)
}

func (ifo *InputFormat) MIMEType() string {
	return C.GoString(ifo.c.mime_type)
}

func (ifo *InputFormat) Extensions() string {
	return C.GoString(ifo.c.extensions)
}

type OutputFormat struct {
	c *C.AVOutputFormat
}

func NewOutputFormat(ptr unsafe.Pointer) *OutputFormat {
	return &OutputFormat{c: (*C.AVOutputFormat)(ptr)}
}

func (of *OutputFormat) Null() bool {
	return of.c == nil
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

func (of *OutputFormat) Flags() int {
	return int(of.c.flags)
}

func (of *OutputFormat) Name() string {
	return C.GoString(of.c.name)
}

func (of *OutputFormat) LongName() string {
	return C.GoString(of.c.long_name)
}

func (of *OutputFormat) MIMEType() string {
	return C.GoString(of.c.mime_type)
}

func (of *OutputFormat) Extensions() string {
	return C.GoString(of.c.extensions)
}

type Stream struct {
	c *C.AVStream
}

func NewStream(ptr unsafe.Pointer) *Stream {
	return &Stream{c: (*C.AVStream)(ptr)}
}

func (s *Stream) Null() bool {
	return s.c == nil
}

func (s *Stream) Unwrap() unsafe.Pointer {
	if s == nil {
		return nil
	}
	return unsafe.Pointer(s.c)
}

func (s *Stream) UnwrapDest() unsafe.Pointer {
	if s == nil {
		return nil
	}
	return unsafe.Pointer(&s.c)
}

func (s *Stream) CodecPar() *avcodec.CodecParameters {
	return avcodec.NewCodecParameters(unsafe.Pointer(s.c.codecpar))
}

func (s *Stream) TimeBase() *avutil.Rational {
	return avutil.NewRational(unsafe.Pointer(&s.c.time_base))
}

func (s *Stream) SetTimeBase(timeBase *avutil.Rational) {
	s.c.time_base = *(*C.AVRational)(timeBase.Unwrap())
}

type FormatContext struct {
	c *C.AVFormatContext
}

func NewFormatContext(ptr unsafe.Pointer) *FormatContext {
	return &FormatContext{c: (*C.AVFormatContext)(ptr)}
}

func (fc *FormatContext) Null() bool {
	return fc.c == nil
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
		formatName0 *C.char
		filename0   = C.CString(filename)
	)
	defer C.free(unsafe.Pointer(filename0))

	if formatName != "" { // zero value
		formatName0 = C.CString(formatName)
		defer C.free(unsafe.Pointer(formatName0))
	}

	return int(C.avformat_alloc_output_context2(&fc.c, (*C.AVOutputFormat)(oformat.Unwrap()), formatName0, filename0))
}

func (fc *FormatContext) OFormat() *OutputFormat {
	return &OutputFormat{c: fc.c.oformat}
}

func (fc *FormatContext) NbStreams() int {
	return int(fc.c.nb_streams)
}

func (fc *FormatContext) Stream(index int) *Stream {
	return &Stream{c: unsafe.Slice(fc.c.streams, int(fc.c.nb_streams))[index]}
}

func (fc *FormatContext) Streams() []*Stream {
	var (
		streamsNum = int(fc.c.nb_streams)

		streams  = make([]*Stream, streamsNum)
		streams0 = unsafe.Slice(fc.c.streams, streamsNum)
	)
	for i, s := range streams0 {
		streams[i] = &Stream{c: s}
	}

	return streams
}

func (fc *FormatContext) NewStream(c *avcodec.Codec) *Stream {
	return &Stream{c: C.avformat_new_stream(fc.c, (*C.AVCodec)(c.Unwrap()))}
}

func (fc *FormatContext) WriteHeader(options *avutil.Dictionary) int {
	return int(C.avformat_write_header(fc.c, (**C.AVDictionary)(options.UnwrapDest())))
}

func (fc *FormatContext) ReadFrame(pkt *avcodec.Packet) int {
	return int(C.av_read_frame(fc.c, (*C.AVPacket)(pkt.Unwrap())))
}

func (fc *FormatContext) InterleavedWriteFrame(pkt *avcodec.Packet) int {
	return int(C.av_interleaved_write_frame(fc.c, (*C.AVPacket)(pkt.Unwrap())))
}

func (fc *FormatContext) WriteTrailer() int {
	return int(C.av_write_trailer(fc.c))
}

func (fc *FormatContext) Pb() *IOContext {
	return &IOContext{c: fc.c.pb}
}

func (fc *FormatContext) SetPb(pb *IOContext) {
	fc.c.pb = pb.c
}

func (fc *FormatContext) Metadata() *avutil.Dictionary {
	return avutil.NewDictionary(unsafe.Pointer(fc.c.metadata))
}

func (fc *FormatContext) FindBestStream(type_ avutil.MediaType, wantedStreamNb int, relatedStream int, flags int) (int, *avcodec.Codec) {
	var decoderRet *C.AVCodec
	return int(C.av_find_best_stream(fc.c, int32(type_), C.int(wantedStreamNb), C.int(relatedStream), &decoderRet, C.int(flags))), avcodec.NewCodec(unsafe.Pointer(decoderRet))
}

func (fc *FormatContext) CloseInput() {
	C.avformat_close_input(&fc.c)
}

func (fc *FormatContext) FreeContext() {
	C.avformat_free_context(fc.c)
}

func MuxerIterate(opaque *ffmpeg.IterationState) *OutputFormat {
	of := C.av_muxer_iterate((*unsafe.Pointer)(opaque.UnwrapDest()))
	if of == nil {
		return nil
	}

	return &OutputFormat{c: of}
}

func DemuxerIterate(opaque *ffmpeg.IterationState) *InputFormat {
	ifo := C.av_demuxer_iterate((*unsafe.Pointer)(opaque.UnwrapDest()))
	if ifo == nil {
		return nil
	}

	return &InputFormat{c: ifo}
}
