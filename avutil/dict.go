package avutil

/*
#cgo pkg-config: libavutil

#include <stdlib.h>
#include <libavutil/dict.h>
*/
import "C"
import "unsafe"

const (
	DictMatchCase     = 1
	DictIgnoreSuffix  = 2
	DictDontStrdupKey = 4
	DictDontStrdupVal = 8
	DictDontOverwrite = 16
	DictAppend        = 32
	DictMultikey      = 64
)

type Dictionary struct {
	c *C.AVDictionary
}

func NewDictionary(ptr unsafe.Pointer) *Dictionary {
	return &Dictionary{c: (*C.AVDictionary)(ptr)}
}

func (d *Dictionary) Null() bool {
	return d.c == nil
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

func (d *Dictionary) Set(key, value string, flags int) int {
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

func (d *Dictionary) Iterate(prev *DictionaryEntry) *DictionaryEntry {
	var prevPtr *C.AVDictionaryEntry
	if prev != nil {
		prevPtr = prev.c
	}

	e := C.av_dict_iterate(d.c, prevPtr)
	if e == nil {
		return nil
	}

	return &DictionaryEntry{c: e}
}

func (d *Dictionary) Free() {
	if d.c != nil {
		C.av_dict_free(&d.c)
	}
}

type DictionaryEntry struct {
	c *C.AVDictionaryEntry
}

func (de *DictionaryEntry) Null() bool {
	return de.c == nil
}

func (de *DictionaryEntry) Unwrap() unsafe.Pointer {
	if de == nil {
		return nil
	}
	return unsafe.Pointer(de.c)
}

func (de *DictionaryEntry) UnwrapDest() unsafe.Pointer {
	if de == nil {
		return nil
	}
	return unsafe.Pointer(&de.c)
}

func (de *DictionaryEntry) Key() string {
	return C.GoString(de.c.key)
}

func (de *DictionaryEntry) Value() string {
	return C.GoString(de.c.value)
}
