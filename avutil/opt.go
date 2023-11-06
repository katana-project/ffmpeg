package avutil

/*
#cgo pkg-config: libavutil

#include <libavutil/opt.h>
*/
import "C"
import "unsafe"

type DictionaryFlag int

const (
	DictionaryFlagMatchCase    DictionaryFlag = iota + 1
	DictionaryFlagIgnoreSuffix DictionaryFlag = 2 << (iota - 1)
	DictionaryFlagDontStrdupKey
	DictionaryFlagDontStrdupVal
	DictionaryFlagDontOverwrite
	DictionaryFlagAppend
	DictionaryFlagMultikey
)

type Dictionary struct {
	c *C.AVDictionary
}

func (d *Dictionary) Unwrap() unsafe.Pointer {
	if d == nil {
		return nil
	}
	return unsafe.Pointer(d.c)
}

func (d *Dictionary) UnwrapDest() unsafe.Pointer {
	if d == nil {
		return nil
	}
	return unsafe.Pointer(&d.c)
}

func (d *Dictionary) Set(key, value string, flags DictionaryFlag) int {
	var (
		key0   = C.CString(key)
		value0 = C.CString(value)
	)
	defer C.free(unsafe.Pointer(key0))
	defer C.free(unsafe.Pointer(value0))

	return int(C.av_dict_set(&d.c, key0, value0, C.int(flags)))
}

func (d *Dictionary) Count() int {
	return int(C.av_dict_count(d.c))
}

func (d *Dictionary) Free() {
	if d.c != nil {
		C.av_dict_free(&d.c)
	}
}
