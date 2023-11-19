package avutil

/*
#cgo pkg-config: libavutil

#include <libavutil/avutil.h>
*/
import "C"

type MediaType int

const (
	MediaTypeUnknown    MediaType = C.AVMEDIA_TYPE_UNKNOWN
	MediaTypeVideo      MediaType = C.AVMEDIA_TYPE_VIDEO
	MediaTypeAudio      MediaType = C.AVMEDIA_TYPE_AUDIO
	MediaTypeData       MediaType = C.AVMEDIA_TYPE_DATA
	MediaTypeSubtitle   MediaType = C.AVMEDIA_TYPE_SUBTITLE
	MediaTypeAttachment MediaType = C.AVMEDIA_TYPE_ATTACHMENT
	MediaTypeNb         MediaType = C.AVMEDIA_TYPE_NB
)

type PictureType int

const (
	PictureTypeNone PictureType = C.AV_PICTURE_TYPE_NONE ///< Undefined
	PictureTypeI    PictureType = C.AV_PICTURE_TYPE_I    ///< Intra
	PictureTypeP    PictureType = C.AV_PICTURE_TYPE_P    ///< Predicted
	PictureTypeB    PictureType = C.AV_PICTURE_TYPE_B    ///< Bi-dir predicted
	PictureTypeS    PictureType = C.AV_PICTURE_TYPE_S    ///< S(GMC)-VOP MPEG-4
	PictureTypeSI   PictureType = C.AV_PICTURE_TYPE_SI   ///< Switching Intra
	PictureTypeSP   PictureType = C.AV_PICTURE_TYPE_SP   ///< Switching Predicted
	PictureTypeBI   PictureType = C.AV_PICTURE_TYPE_BI   ///< BI type
)
