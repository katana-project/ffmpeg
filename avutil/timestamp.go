package avutil

/*
#cgo pkg-config: libavutil

#include <libavutil/timestamp.h>

char* av_ts2str_heap(int64_t ts) {
  char* buffer = malloc(AV_TS_MAX_STRING_SIZE);
  if (buffer == NULL) {
    return NULL;
  }

  return av_ts_make_string(buffer, ts);
}

char* av_ts2timestr_heap(int64_t ts, AVRational* tb) {
  char* buffer = malloc(AV_TS_MAX_STRING_SIZE);
  if (buffer == NULL) {
    return NULL;
  }

  return av_ts_make_time_string(buffer, ts, tb);
}
*/
import "C"
import "unsafe"

func Ts2Str(ts int64) string {
	tsStr := C.av_ts2str_heap(C.int64_t(ts))
	if tsStr == nil {
		panic("memory allocation failed")
	}

	defer C.free(unsafe.Pointer(tsStr))
	return C.GoString(tsStr)
}

func Ts2TimeStr(ts int64, tb *Rational) string {
	timeStr := C.av_ts2timestr_heap(C.int64_t(ts), &tb.c)
	if timeStr == nil {
		panic("memory allocation failed")
	}

	defer C.free(unsafe.Pointer(timeStr))
	return C.GoString(timeStr)
}
