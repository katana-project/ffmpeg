package avutil

/*
#cgo pkg-config: libavutil

#include <libavutil/log.h>
*/
import "C"

const (
	LogQuiet   = -8
	LogPanic   = 0
	LogFatal   = 8
	LogError   = 16
	LogWarning = 24
	LogInfo    = 32
	LogVerbose = 40
	LogDebug   = 48
	LogTrace   = 56
)

func LogLevel() int {
	return int(C.av_log_get_level())
}

func SetLogLevel(level int) {
	C.av_log_set_level(C.int(level))
}
