package avutil

/*
#cgo pkg-config: libavutil

#include <stdlib.h>
#include <libavutil/error.h>

const int macro_AVERROR_EOF = AVERROR_EOF;
const int macro_AVERROR_UNKNOWN = AVERROR_UNKNOWN;
const int macro_AVERROR_ENOMEM = AVERROR(ENOMEM);

char* av_err2str_heap(int errnum) {
  char* buffer = malloc(AV_ERROR_MAX_STRING_SIZE);
  if (buffer == NULL) {
    return NULL;
  }

  return av_make_error_string(buffer, AV_ERROR_MAX_STRING_SIZE, errnum);
}
*/
import "C"
import "unsafe"

var (
	ErrorEOF     = int(C.macro_AVERROR_EOF)
	ErrorUnknown = int(C.macro_AVERROR_UNKNOWN)
	ErrorENoMem  = int(C.macro_AVERROR_ENOMEM)
)

func Err2Str(errnum int) string {
	errStr := C.av_err2str_heap(C.int(errnum))
	if errStr == nil {
		panic("memory allocation failed")
	}

	defer C.free(unsafe.Pointer(errStr))
	return C.GoString(errStr)
}
