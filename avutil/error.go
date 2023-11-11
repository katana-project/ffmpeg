package avutil

/*
#cgo pkg-config: libavutil

#include <stdlib.h>
#include <libavutil/error.h>

int macro_AVERROR_EOF() {
  return AVERROR_EOF;
}

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
	ErrorEOF = int(C.macro_AVERROR_EOF())
)

func Err2Str(errnum int) string {
	errStr := C.av_err2str_heap(C.int(errnum))
	defer C.free(unsafe.Pointer(errStr))

	return C.GoString(errStr)
}
