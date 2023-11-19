package avutil

/*
#cgo pkg-config: libavutil

#include <libavutil/pixfmt.h>
*/
import "C"

type PixelFormat int

const (
	PixelFormatNONE      PixelFormat = C.AV_PIX_FMT_NONE
	PixelFormatYUV420P   PixelFormat = C.AV_PIX_FMT_YUV420P   ///< planar YUV 4:2:0, 12bpp, (1 Cr & Cb sample per 2x2 Y samples)
	PixelFormatYUYV422   PixelFormat = C.AV_PIX_FMT_YUYV422   ///< packed YUV 4:2:2, 16bpp, Y0 Cb Y1 Cr
	PixelFormatRGB24     PixelFormat = C.AV_PIX_FMT_RGB24     ///< packed RGB 8:8:8, 24bpp, RGBRGB...
	PixelFormatBGR24     PixelFormat = C.AV_PIX_FMT_BGR24     ///< packed RGB 8:8:8, 24bpp, BGRBGR...
	PixelFormatYUV422P   PixelFormat = C.AV_PIX_FMT_YUV422P   ///< planar YUV 4:2:2, 16bpp, (1 Cr & Cb sample per 2x1 Y samples)
	PixelFormatYUV444P   PixelFormat = C.AV_PIX_FMT_YUV444P   ///< planar YUV 4:4:4, 24bpp, (1 Cr & Cb sample per 1x1 Y samples)
	PixelFormatYUV410P   PixelFormat = C.AV_PIX_FMT_YUV410P   ///< planar YUV 4:1:0,  9bpp, (1 Cr & Cb sample per 4x4 Y samples)
	PixelFormatYUV411P   PixelFormat = C.AV_PIX_FMT_YUV411P   ///< planar YUV 4:1:1, 12bpp, (1 Cr & Cb sample per 4x1 Y samples)
	PixelFormatGRAY8     PixelFormat = C.AV_PIX_FMT_GRAY8     ///<        Y        ,  8bpp
	PixelFormatMONOWHITE PixelFormat = C.AV_PIX_FMT_MONOWHITE ///<        Y        ,  1bpp, 0 is white, 1 is black, in each byte pixels are ordered from the msb to the lsb
	PixelFormatMONOBLACK PixelFormat = C.AV_PIX_FMT_MONOBLACK ///<        Y        ,  1bpp, 0 is black, 1 is white, in each byte pixels are ordered from the msb to the lsb
	PixelFormatPAL8      PixelFormat = C.AV_PIX_FMT_PAL8      ///< 8 bits with AV_PIX_FMT_RGB32 palette
	PixelFormatYUVJ420P  PixelFormat = C.AV_PIX_FMT_YUVJ420P  ///< planar YUV 4:2:0, 12bpp, full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV420P and setting color_range
	PixelFormatYUVJ422P  PixelFormat = C.AV_PIX_FMT_YUVJ422P  ///< planar YUV 4:2:2, 16bpp, full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV422P and setting color_range
	PixelFormatYUVJ444P  PixelFormat = C.AV_PIX_FMT_YUVJ444P  ///< planar YUV 4:4:4, 24bpp, full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV444P and setting color_range
	PixelFormatUYVY422   PixelFormat = C.AV_PIX_FMT_UYVY422   ///< packed YUV 4:2:2, 16bpp, Cb Y0 Cr Y1
	PixelFormatUYYVYY411 PixelFormat = C.AV_PIX_FMT_UYYVYY411 ///< packed YUV 4:1:1, 12bpp, Cb Y0 Y1 Cr Y2 Y3
	PixelFormatBGR8      PixelFormat = C.AV_PIX_FMT_BGR8      ///< packed RGB 3:3:2,  8bpp, (msb)2B 3G 3R(lsb)
	PixelFormatBGR4      PixelFormat = C.AV_PIX_FMT_BGR4      ///< packed RGB 1:2:1 bitstream,  4bpp, (msb)1B 2G 1R(lsb), a byte contains two pixels, the first pixel in the byte is the one composed by the 4 msb bits
	PixelFormatBGR4_BYTE PixelFormat = C.AV_PIX_FMT_BGR4_BYTE ///< packed RGB 1:2:1,  8bpp, (msb)1B 2G 1R(lsb)
	PixelFormatRGB8      PixelFormat = C.AV_PIX_FMT_RGB8      ///< packed RGB 3:3:2,  8bpp, (msb)2R 3G 3B(lsb)
	PixelFormatRGB4      PixelFormat = C.AV_PIX_FMT_RGB4      ///< packed RGB 1:2:1 bitstream,  4bpp, (msb)1R 2G 1B(lsb), a byte contains two pixels, the first pixel in the byte is the one composed by the 4 msb bits
	PixelFormatRGB4_BYTE PixelFormat = C.AV_PIX_FMT_RGB4_BYTE ///< packed RGB 1:2:1,  8bpp, (msb)1R 2G 1B(lsb)
	PixelFormatNV12      PixelFormat = C.AV_PIX_FMT_NV12      ///< planar YUV 4:2:0, 12bpp, 1 plane for Y and 1 plane for the UV components, which are interleaved (first byte U and the following byte V)
	PixelFormatNV21      PixelFormat = C.AV_PIX_FMT_NV21      ///< as above, but U and V bytes are swapped

	PixelFormatARGB PixelFormat = C.AV_PIX_FMT_ARGB ///< packed ARGB 8:8:8:8, 32bpp, ARGBARGB...
	PixelFormatRGBA PixelFormat = C.AV_PIX_FMT_RGBA ///< packed RGBA 8:8:8:8, 32bpp, RGBARGBA...
	PixelFormatABGR PixelFormat = C.AV_PIX_FMT_ABGR ///< packed ABGR 8:8:8:8, 32bpp, ABGRABGR...
	PixelFormatBGRA PixelFormat = C.AV_PIX_FMT_BGRA ///< packed BGRA 8:8:8:8, 32bpp, BGRABGRA...

	PixelFormatGRAY16BE PixelFormat = C.AV_PIX_FMT_GRAY16BE ///<        Y        , 16bpp, big-endian
	PixelFormatGRAY16LE PixelFormat = C.AV_PIX_FMT_GRAY16LE ///<        Y        , 16bpp, little-endian
	PixelFormatYUV440P  PixelFormat = C.AV_PIX_FMT_YUV440P  ///< planar YUV 4:4:0 (1 Cr & Cb sample per 1x2 Y samples)
	PixelFormatYUVJ440P PixelFormat = C.AV_PIX_FMT_YUVJ440P ///< planar YUV 4:4:0 full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV440P and setting color_range
	PixelFormatYUVA420P PixelFormat = C.AV_PIX_FMT_YUVA420P ///< planar YUV 4:2:0, 20bpp, (1 Cr & Cb sample per 2x2 Y & A samples)
	PixelFormatRGB48BE  PixelFormat = C.AV_PIX_FMT_RGB48BE  ///< packed RGB 16:16:16, 48bpp, 16R, 16G, 16B, the 2-byte value for each R/G/B component is stored as big-endian
	PixelFormatRGB48LE  PixelFormat = C.AV_PIX_FMT_RGB48LE  ///< packed RGB 16:16:16, 48bpp, 16R, 16G, 16B, the 2-byte value for each R/G/B component is stored as little-endian

	PixelFormatRGB565BE PixelFormat = C.AV_PIX_FMT_RGB565BE ///< packed RGB 5:6:5, 16bpp, (msb)   5R 6G 5B(lsb), big-endian
	PixelFormatRGB565LE PixelFormat = C.AV_PIX_FMT_RGB565LE ///< packed RGB 5:6:5, 16bpp, (msb)   5R 6G 5B(lsb), little-endian
	PixelFormatRGB555BE PixelFormat = C.AV_PIX_FMT_RGB555BE ///< packed RGB 5:5:5, 16bpp, (msb)1X 5R 5G 5B(lsb), big-endian   , X=unused/undefined
	PixelFormatRGB555LE PixelFormat = C.AV_PIX_FMT_RGB555LE ///< packed RGB 5:5:5, 16bpp, (msb)1X 5R 5G 5B(lsb), little-endian, X=unused/undefined

	PixelFormatBGR565BE PixelFormat = C.AV_PIX_FMT_BGR565BE ///< packed BGR 5:6:5, 16bpp, (msb)   5B 6G 5R(lsb), big-endian
	PixelFormatBGR565LE PixelFormat = C.AV_PIX_FMT_BGR565LE ///< packed BGR 5:6:5, 16bpp, (msb)   5B 6G 5R(lsb), little-endian
	PixelFormatBGR555BE PixelFormat = C.AV_PIX_FMT_BGR555BE ///< packed BGR 5:5:5, 16bpp, (msb)1X 5B 5G 5R(lsb), big-endian   , X=unused/undefined
	PixelFormatBGR555LE PixelFormat = C.AV_PIX_FMT_BGR555LE ///< packed BGR 5:5:5, 16bpp, (msb)1X 5B 5G 5R(lsb), little-endian, X=unused/undefined

	// Hardware acceleration through VA-API, data[3] contains a VASurfaceID.

	PixelFormatVAAPI PixelFormat = C.AV_PIX_FMT_VAAPI

	PixelFormatYUV420P16LE PixelFormat = C.AV_PIX_FMT_YUV420P16LE ///< planar YUV 4:2:0, 24bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	PixelFormatYUV420P16BE PixelFormat = C.AV_PIX_FMT_YUV420P16BE ///< planar YUV 4:2:0, 24bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	PixelFormatYUV422P16LE PixelFormat = C.AV_PIX_FMT_YUV422P16LE ///< planar YUV 4:2:2, 32bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	PixelFormatYUV422P16BE PixelFormat = C.AV_PIX_FMT_YUV422P16BE ///< planar YUV 4:2:2, 32bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	PixelFormatYUV444P16LE PixelFormat = C.AV_PIX_FMT_YUV444P16LE ///< planar YUV 4:4:4, 48bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	PixelFormatYUV444P16BE PixelFormat = C.AV_PIX_FMT_YUV444P16BE ///< planar YUV 4:4:4, 48bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	PixelFormatDXVA2_VLD   PixelFormat = C.AV_PIX_FMT_DXVA2_VLD   ///< HW decoding through DXVA2, Picture.data[3] contains a LPDIRECT3DSURFACE9 pointer

	PixelFormatRGB444LE PixelFormat = C.AV_PIX_FMT_RGB444LE ///< packed RGB 4:4:4, 16bpp, (msb)4X 4R 4G 4B(lsb), little-endian, X=unused/undefined
	PixelFormatRGB444BE PixelFormat = C.AV_PIX_FMT_RGB444BE ///< packed RGB 4:4:4, 16bpp, (msb)4X 4R 4G 4B(lsb), big-endian,    X=unused/undefined
	PixelFormatBGR444LE PixelFormat = C.AV_PIX_FMT_BGR444LE ///< packed BGR 4:4:4, 16bpp, (msb)4X 4B 4G 4R(lsb), little-endian, X=unused/undefined
	PixelFormatBGR444BE PixelFormat = C.AV_PIX_FMT_BGR444BE ///< packed BGR 4:4:4, 16bpp, (msb)4X 4B 4G 4R(lsb), big-endian,    X=unused/undefined
	PixelFormatYA8      PixelFormat = C.AV_PIX_FMT_YA8      ///< 8 bits gray, 8 bits alpha, alias AV_PIX_FMT_Y400A and AV_PIX_FMT_GRAY8A
	PixelFormatBGR48BE  PixelFormat = C.AV_PIX_FMT_BGR48BE  ///< packed RGB 16:16:16, 48bpp, 16B, 16G, 16R, the 2-byte value for each R/G/B component is stored as big-endian
	PixelFormatBGR48LE  PixelFormat = C.AV_PIX_FMT_BGR48LE  ///< packed RGB 16:16:16, 48bpp, 16B, 16G, 16R, the 2-byte value for each R/G/B component is stored as little-endian

	// The following 12 formats have the disadvantage of needing 1 format for each bit depth.
	// Notice that each 9/10 bits sample is stored in 16 bits with extra padding.
	// If you want to support multiple bit depths, then using AV_PIX_FMT_YUV420P16* with the bpp stored separately is better.

	PixelFormatYUV420P9BE   PixelFormat = C.AV_PIX_FMT_YUV420P9BE   ///< planar YUV 4:2:0, 13.5bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	PixelFormatYUV420P9LE   PixelFormat = C.AV_PIX_FMT_YUV420P9LE   ///< planar YUV 4:2:0, 13.5bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	PixelFormatYUV420P10BE  PixelFormat = C.AV_PIX_FMT_YUV420P10BE  ///< planar YUV 4:2:0, 15bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	PixelFormatYUV420P10LE  PixelFormat = C.AV_PIX_FMT_YUV420P10LE  ///< planar YUV 4:2:0, 15bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	PixelFormatYUV422P10BE  PixelFormat = C.AV_PIX_FMT_YUV422P10BE  ///< planar YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	PixelFormatYUV422P10LE  PixelFormat = C.AV_PIX_FMT_YUV422P10LE  ///< planar YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	PixelFormatYUV444P9BE   PixelFormat = C.AV_PIX_FMT_YUV444P9BE   ///< planar YUV 4:4:4, 27bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	PixelFormatYUV444P9LE   PixelFormat = C.AV_PIX_FMT_YUV444P9LE   ///< planar YUV 4:4:4, 27bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	PixelFormatYUV444P10BE  PixelFormat = C.AV_PIX_FMT_YUV444P10BE  ///< planar YUV 4:4:4, 30bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	PixelFormatYUV444P10LE  PixelFormat = C.AV_PIX_FMT_YUV444P10LE  ///< planar YUV 4:4:4, 30bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	PixelFormatYUV422P9BE   PixelFormat = C.AV_PIX_FMT_YUV422P9BE   ///< planar YUV 4:2:2, 18bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	PixelFormatYUV422P9LE   PixelFormat = C.AV_PIX_FMT_YUV422P9LE   ///< planar YUV 4:2:2, 18bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	PixelFormatGBRP         PixelFormat = C.AV_PIX_FMT_GBRP         ///< planar GBR 4:4:4 24bpp, alias AV_PIX_FMT_GBR24P
	PixelFormatGBRP9BE      PixelFormat = C.AV_PIX_FMT_GBRP9BE      ///< planar GBR 4:4:4 27bpp, big-endian
	PixelFormatGBRP9LE      PixelFormat = C.AV_PIX_FMT_GBRP9LE      ///< planar GBR 4:4:4 27bpp, little-endian
	PixelFormatGBRP10BE     PixelFormat = C.AV_PIX_FMT_GBRP10BE     ///< planar GBR 4:4:4 30bpp, big-endian
	PixelFormatGBRP10LE     PixelFormat = C.AV_PIX_FMT_GBRP10LE     ///< planar GBR 4:4:4 30bpp, little-endian
	PixelFormatGBRP16BE     PixelFormat = C.AV_PIX_FMT_GBRP16BE     ///< planar GBR 4:4:4 48bpp, big-endian
	PixelFormatGBRP16LE     PixelFormat = C.AV_PIX_FMT_GBRP16LE     ///< planar GBR 4:4:4 48bpp, little-endian
	PixelFormatYUVA422P     PixelFormat = C.AV_PIX_FMT_YUVA422P     ///< planar YUV 4:2:2 24bpp, (1 Cr & Cb sample per 2x1 Y & A samples)
	PixelFormatYUVA444P     PixelFormat = C.AV_PIX_FMT_YUVA444P     ///< planar YUV 4:4:4 32bpp, (1 Cr & Cb sample per 1x1 Y & A samples)
	PixelFormatYUVA420P9BE  PixelFormat = C.AV_PIX_FMT_YUVA420P9BE  ///< planar YUV 4:2:0 22.5bpp, (1 Cr & Cb sample per 2x2 Y & A samples), big-endian
	PixelFormatYUVA420P9LE  PixelFormat = C.AV_PIX_FMT_YUVA420P9LE  ///< planar YUV 4:2:0 22.5bpp, (1 Cr & Cb sample per 2x2 Y & A samples), little-endian
	PixelFormatYUVA422P9BE  PixelFormat = C.AV_PIX_FMT_YUVA422P9BE  ///< planar YUV 4:2:2 27bpp, (1 Cr & Cb sample per 2x1 Y & A samples), big-endian
	PixelFormatYUVA422P9LE  PixelFormat = C.AV_PIX_FMT_YUVA422P9LE  ///< planar YUV 4:2:2 27bpp, (1 Cr & Cb sample per 2x1 Y & A samples), little-endian
	PixelFormatYUVA444P9BE  PixelFormat = C.AV_PIX_FMT_YUVA444P9BE  ///< planar YUV 4:4:4 36bpp, (1 Cr & Cb sample per 1x1 Y & A samples), big-endian
	PixelFormatYUVA444P9LE  PixelFormat = C.AV_PIX_FMT_YUVA444P9LE  ///< planar YUV 4:4:4 36bpp, (1 Cr & Cb sample per 1x1 Y & A samples), little-endian
	PixelFormatYUVA420P10BE PixelFormat = C.AV_PIX_FMT_YUVA420P10BE ///< planar YUV 4:2:0 25bpp, (1 Cr & Cb sample per 2x2 Y & A samples, big-endian)
	PixelFormatYUVA420P10LE PixelFormat = C.AV_PIX_FMT_YUVA420P10LE ///< planar YUV 4:2:0 25bpp, (1 Cr & Cb sample per 2x2 Y & A samples, little-endian)
	PixelFormatYUVA422P10BE PixelFormat = C.AV_PIX_FMT_YUVA422P10BE ///< planar YUV 4:2:2 30bpp, (1 Cr & Cb sample per 2x1 Y & A samples, big-endian)
	PixelFormatYUVA422P10LE PixelFormat = C.AV_PIX_FMT_YUVA422P10LE ///< planar YUV 4:2:2 30bpp, (1 Cr & Cb sample per 2x1 Y & A samples, little-endian)
	PixelFormatYUVA444P10BE PixelFormat = C.AV_PIX_FMT_YUVA444P10BE ///< planar YUV 4:4:4 40bpp, (1 Cr & Cb sample per 1x1 Y & A samples, big-endian)
	PixelFormatYUVA444P10LE PixelFormat = C.AV_PIX_FMT_YUVA444P10LE ///< planar YUV 4:4:4 40bpp, (1 Cr & Cb sample per 1x1 Y & A samples, little-endian)
	PixelFormatYUVA420P16BE PixelFormat = C.AV_PIX_FMT_YUVA420P16BE ///< planar YUV 4:2:0 40bpp, (1 Cr & Cb sample per 2x2 Y & A samples, big-endian)
	PixelFormatYUVA420P16LE PixelFormat = C.AV_PIX_FMT_YUVA420P16LE ///< planar YUV 4:2:0 40bpp, (1 Cr & Cb sample per 2x2 Y & A samples, little-endian)
	PixelFormatYUVA422P16BE PixelFormat = C.AV_PIX_FMT_YUVA422P16BE ///< planar YUV 4:2:2 48bpp, (1 Cr & Cb sample per 2x1 Y & A samples, big-endian)
	PixelFormatYUVA422P16LE PixelFormat = C.AV_PIX_FMT_YUVA422P16LE ///< planar YUV 4:2:2 48bpp, (1 Cr & Cb sample per 2x1 Y & A samples, little-endian)
	PixelFormatYUVA444P16BE PixelFormat = C.AV_PIX_FMT_YUVA444P16BE ///< planar YUV 4:4:4 64bpp, (1 Cr & Cb sample per 1x1 Y & A samples, big-endian)
	PixelFormatYUVA444P16LE PixelFormat = C.AV_PIX_FMT_YUVA444P16LE ///< planar YUV 4:4:4 64bpp, (1 Cr & Cb sample per 1x1 Y & A samples, little-endian)

	PixelFormatVDPAU PixelFormat = C.AV_PIX_FMT_VDPAU ///< HW acceleration through VDPAU, Picture.data[3] contains a VdpVideoSurface

	PixelFormatXYZ12LE PixelFormat = C.AV_PIX_FMT_XYZ12LE ///< packed XYZ 4:4:4, 36 bpp, (msb) 12X, 12Y, 12Z (lsb), the 2-byte value for each X/Y/Z is stored as little-endian, the 4 lower bits are set to 0
	PixelFormatXYZ12BE PixelFormat = C.AV_PIX_FMT_XYZ12BE ///< packed XYZ 4:4:4, 36 bpp, (msb) 12X, 12Y, 12Z (lsb), the 2-byte value for each X/Y/Z is stored as big-endian, the 4 lower bits are set to 0
	PixelFormatNV16    PixelFormat = C.AV_PIX_FMT_NV16    ///< interleaved chroma YUV 4:2:2, 16bpp, (1 Cr & Cb sample per 2x1 Y samples)
	PixelFormatNV20LE  PixelFormat = C.AV_PIX_FMT_NV20LE  ///< interleaved chroma YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	PixelFormatNV20BE  PixelFormat = C.AV_PIX_FMT_NV20BE  ///< interleaved chroma YUV 4:2:2, 20bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian

	PixelFormatRGBA64BE PixelFormat = C.AV_PIX_FMT_RGBA64BE ///< packed RGBA 16:16:16:16, 64bpp, 16R, 16G, 16B, 16A, the 2-byte value for each R/G/B/A component is stored as big-endian
	PixelFormatRGBA64LE PixelFormat = C.AV_PIX_FMT_RGBA64LE ///< packed RGBA 16:16:16:16, 64bpp, 16R, 16G, 16B, 16A, the 2-byte value for each R/G/B/A component is stored as little-endian
	PixelFormatBGRA64BE PixelFormat = C.AV_PIX_FMT_BGRA64BE ///< packed RGBA 16:16:16:16, 64bpp, 16B, 16G, 16R, 16A, the 2-byte value for each R/G/B/A component is stored as big-endian
	PixelFormatBGRA64LE PixelFormat = C.AV_PIX_FMT_BGRA64LE ///< packed RGBA 16:16:16:16, 64bpp, 16B, 16G, 16R, 16A, the 2-byte value for each R/G/B/A component is stored as little-endian

	PixelFormatYVYU422 PixelFormat = C.AV_PIX_FMT_YVYU422 ///< packed YUV 4:2:2, 16bpp, Y0 Cr Y1 Cb

	PixelFormatYA16BE PixelFormat = C.AV_PIX_FMT_YA16BE ///< 16 bits gray, 16 bits alpha (big-endian)
	PixelFormatYA16LE PixelFormat = C.AV_PIX_FMT_YA16LE ///< 16 bits gray, 16 bits alpha (little-endian)

	PixelFormatGBRAP     PixelFormat = C.AV_PIX_FMT_GBRAP     ///< planar GBRA 4:4:4:4 32bpp
	PixelFormatGBRAP16BE PixelFormat = C.AV_PIX_FMT_GBRAP16BE ///< planar GBRA 4:4:4:4 64bpp, big-endian
	PixelFormatGBRAP16LE PixelFormat = C.AV_PIX_FMT_GBRAP16LE ///< planar GBRA 4:4:4:4 64bpp, little-endian

	// HW acceleration through QSV, data[3] contains a pointer to the
	// mfxFrameSurface1 structure.
	//
	// Before FFmpeg 5.0:
	// mfxFrameSurface1.Data.MemId contains a pointer when importing
	// the following frames as QSV frames:
	//
	// VAAPI:
	// mfxFrameSurface1.Data.MemId contains a pointer to VASurfaceID
	//
	// DXVA2:
	// mfxFrameSurface1.Data.MemId contains a pointer to IDirect3DSurface9
	//
	// FFmpeg 5.0 and above:
	// mfxFrameSurface1.Data.MemId contains a pointer to the mfxHDLPair
	// structure when importing the following frames as QSV frames:
	//
	// VAAPI:
	// mfxHDLPair.first contains a VASurfaceID pointer.
	// mfxHDLPair.second is always MFX_INFINITE.
	//
	// DXVA2:
	// mfxHDLPair.first contains IDirect3DSurface9 pointer.
	// mfxHDLPair.second is always MFX_INFINITE.
	//
	// D3D11:
	// mfxHDLPair.first contains a ID3D11Texture2D pointer.
	// mfxHDLPair.second contains the texture array index of the frame if the
	// ID3D11Texture2D is an array texture, or always MFX_INFINITE if it is a
	// normal texture.

	PixelFormatQSV PixelFormat = C.AV_PIX_FMT_QSV

	// HW acceleration though MMAL, data[3] contains a pointer to the MMAL_BUFFER_HEADER_T structure.

	PixelFormatMMAL PixelFormat = C.AV_PIX_FMT_MMAL

	PixelFormatD3D11VA_VLD PixelFormat = C.AV_PIX_FMT_D3D11VA_VLD ///< HW decoding through Direct3D11 via old API, Picture.data[3] contains a ID3D11VideoDecoderOutputView pointer

	// HW acceleration through CUDA. data[i] contain CUdeviceptr pointers exactly as for system memory frames.

	PixelFormatCUDA PixelFormat = C.AV_PIX_FMT_CUDA

	PixelFormat0RGB PixelFormat = C.AV_PIX_FMT_0RGB ///< packed RGB 8:8:8, 32bpp, XRGBXRGB...   X=unused/undefined
	PixelFormatRGB0 PixelFormat = C.AV_PIX_FMT_RGB0 ///< packed RGB 8:8:8, 32bpp, RGBXRGBX...   X=unused/undefined
	PixelFormat0BGR PixelFormat = C.AV_PIX_FMT_0BGR ///< packed BGR 8:8:8, 32bpp, XBGRXBGR...   X=unused/undefined
	PixelFormatBGR0 PixelFormat = C.AV_PIX_FMT_BGR0 ///< packed BGR 8:8:8, 32bpp, BGRXBGRX...   X=unused/undefined

	PixelFormatYUV420P12BE PixelFormat = C.AV_PIX_FMT_YUV420P12BE ///< planar YUV 4:2:0,18bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	PixelFormatYUV420P12LE PixelFormat = C.AV_PIX_FMT_YUV420P12LE ///< planar YUV 4:2:0,18bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	PixelFormatYUV420P14BE PixelFormat = C.AV_PIX_FMT_YUV420P14BE ///< planar YUV 4:2:0,21bpp, (1 Cr & Cb sample per 2x2 Y samples), big-endian
	PixelFormatYUV420P14LE PixelFormat = C.AV_PIX_FMT_YUV420P14LE ///< planar YUV 4:2:0,21bpp, (1 Cr & Cb sample per 2x2 Y samples), little-endian
	PixelFormatYUV422P12BE PixelFormat = C.AV_PIX_FMT_YUV422P12BE ///< planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	PixelFormatYUV422P12LE PixelFormat = C.AV_PIX_FMT_YUV422P12LE ///< planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	PixelFormatYUV422P14BE PixelFormat = C.AV_PIX_FMT_YUV422P14BE ///< planar YUV 4:2:2,28bpp, (1 Cr & Cb sample per 2x1 Y samples), big-endian
	PixelFormatYUV422P14LE PixelFormat = C.AV_PIX_FMT_YUV422P14LE ///< planar YUV 4:2:2,28bpp, (1 Cr & Cb sample per 2x1 Y samples), little-endian
	PixelFormatYUV444P12BE PixelFormat = C.AV_PIX_FMT_YUV444P12BE ///< planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	PixelFormatYUV444P12LE PixelFormat = C.AV_PIX_FMT_YUV444P12LE ///< planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	PixelFormatYUV444P14BE PixelFormat = C.AV_PIX_FMT_YUV444P14BE ///< planar YUV 4:4:4,42bpp, (1 Cr & Cb sample per 1x1 Y samples), big-endian
	PixelFormatYUV444P14LE PixelFormat = C.AV_PIX_FMT_YUV444P14LE ///< planar YUV 4:4:4,42bpp, (1 Cr & Cb sample per 1x1 Y samples), little-endian
	PixelFormatGBRP12BE    PixelFormat = C.AV_PIX_FMT_GBRP12BE    ///< planar GBR 4:4:4 36bpp, big-endian
	PixelFormatGBRP12LE    PixelFormat = C.AV_PIX_FMT_GBRP12LE    ///< planar GBR 4:4:4 36bpp, little-endian
	PixelFormatGBRP14BE    PixelFormat = C.AV_PIX_FMT_GBRP14BE    ///< planar GBR 4:4:4 42bpp, big-endian
	PixelFormatGBRP14LE    PixelFormat = C.AV_PIX_FMT_GBRP14LE    ///< planar GBR 4:4:4 42bpp, little-endian
	PixelFormatYUVJ411P    PixelFormat = C.AV_PIX_FMT_YUVJ411P    ///< planar YUV 4:1:1, 12bpp, (1 Cr & Cb sample per 4x1 Y samples) full scale (JPEG), deprecated in favor of AV_PIX_FMT_YUV411P and setting color_range

	PixelFormatBAYER_BGGR8    PixelFormat = C.AV_PIX_FMT_BAYER_BGGR8    ///< bayer, BGBG..(odd line), GRGR..(even line), 8-bit samples
	PixelFormatBAYER_RGGB8    PixelFormat = C.AV_PIX_FMT_BAYER_RGGB8    ///< bayer, RGRG..(odd line), GBGB..(even line), 8-bit samples
	PixelFormatBAYER_GBRG8    PixelFormat = C.AV_PIX_FMT_BAYER_GBRG8    ///< bayer, GBGB..(odd line), RGRG..(even line), 8-bit samples
	PixelFormatBAYER_GRBG8    PixelFormat = C.AV_PIX_FMT_BAYER_GRBG8    ///< bayer, GRGR..(odd line), BGBG..(even line), 8-bit samples
	PixelFormatBAYER_BGGR16LE PixelFormat = C.AV_PIX_FMT_BAYER_BGGR16LE ///< bayer, BGBG..(odd line), GRGR..(even line), 16-bit samples, little-endian
	PixelFormatBAYER_BGGR16BE PixelFormat = C.AV_PIX_FMT_BAYER_BGGR16BE ///< bayer, BGBG..(odd line), GRGR..(even line), 16-bit samples, big-endian
	PixelFormatBAYER_RGGB16LE PixelFormat = C.AV_PIX_FMT_BAYER_RGGB16LE ///< bayer, RGRG..(odd line), GBGB..(even line), 16-bit samples, little-endian
	PixelFormatBAYER_RGGB16BE PixelFormat = C.AV_PIX_FMT_BAYER_RGGB16BE ///< bayer, RGRG..(odd line), GBGB..(even line), 16-bit samples, big-endian
	PixelFormatBAYER_GBRG16LE PixelFormat = C.AV_PIX_FMT_BAYER_GBRG16LE ///< bayer, GBGB..(odd line), RGRG..(even line), 16-bit samples, little-endian
	PixelFormatBAYER_GBRG16BE PixelFormat = C.AV_PIX_FMT_BAYER_GBRG16BE ///< bayer, GBGB..(odd line), RGRG..(even line), 16-bit samples, big-endian
	PixelFormatBAYER_GRBG16LE PixelFormat = C.AV_PIX_FMT_BAYER_GRBG16LE ///< bayer, GRGR..(odd line), BGBG..(even line), 16-bit samples, little-endian
	PixelFormatBAYER_GRBG16BE PixelFormat = C.AV_PIX_FMT_BAYER_GRBG16BE ///< bayer, GRGR..(odd line), BGBG..(even line), 16-bit samples, big-endian

	/*#if FF_API_XVMC
	  PixelFormatXVMC PixelFormat = C.AV_PIX_FMT_XVMC///< XVideo Motion Acceleration via common packet passing
	  #endif*/

	PixelFormatYUV440P10LE PixelFormat = C.AV_PIX_FMT_YUV440P10LE ///< planar YUV 4:4:0,20bpp, (1 Cr & Cb sample per 1x2 Y samples), little-endian
	PixelFormatYUV440P10BE PixelFormat = C.AV_PIX_FMT_YUV440P10BE ///< planar YUV 4:4:0,20bpp, (1 Cr & Cb sample per 1x2 Y samples), big-endian
	PixelFormatYUV440P12LE PixelFormat = C.AV_PIX_FMT_YUV440P12LE ///< planar YUV 4:4:0,24bpp, (1 Cr & Cb sample per 1x2 Y samples), little-endian
	PixelFormatYUV440P12BE PixelFormat = C.AV_PIX_FMT_YUV440P12BE ///< planar YUV 4:4:0,24bpp, (1 Cr & Cb sample per 1x2 Y samples), big-endian
	PixelFormatAYUV64LE    PixelFormat = C.AV_PIX_FMT_AYUV64LE    ///< packed AYUV 4:4:4,64bpp (1 Cr & Cb sample per 1x1 Y & A samples), little-endian
	PixelFormatAYUV64BE    PixelFormat = C.AV_PIX_FMT_AYUV64BE    ///< packed AYUV 4:4:4,64bpp (1 Cr & Cb sample per 1x1 Y & A samples), big-endian

	PixelFormatVIDEOTOOLBOX PixelFormat = C.AV_PIX_FMT_VIDEOTOOLBOX ///< hardware decoding through Videotoolbox

	PixelFormatP010LE PixelFormat = C.AV_PIX_FMT_P010LE ///< like NV12, with 10bpp per component, data in the high bits, zeros in the low bits, little-endian
	PixelFormatP010BE PixelFormat = C.AV_PIX_FMT_P010BE ///< like NV12, with 10bpp per component, data in the high bits, zeros in the low bits, big-endian

	PixelFormatGBRAP12BE PixelFormat = C.AV_PIX_FMT_GBRAP12BE ///< planar GBR 4:4:4:4 48bpp, big-endian
	PixelFormatGBRAP12LE PixelFormat = C.AV_PIX_FMT_GBRAP12LE ///< planar GBR 4:4:4:4 48bpp, little-endian

	PixelFormatGBRAP10BE PixelFormat = C.AV_PIX_FMT_GBRAP10BE ///< planar GBR 4:4:4:4 40bpp, big-endian
	PixelFormatGBRAP10LE PixelFormat = C.AV_PIX_FMT_GBRAP10LE ///< planar GBR 4:4:4:4 40bpp, little-endian

	PixelFormatMEDIACODEC PixelFormat = C.AV_PIX_FMT_MEDIACODEC ///< hardware decoding through MediaCodec

	PixelFormatGRAY12BE PixelFormat = C.AV_PIX_FMT_GRAY12BE ///<        Y        , 12bpp, big-endian
	PixelFormatGRAY12LE PixelFormat = C.AV_PIX_FMT_GRAY12LE ///<        Y        , 12bpp, little-endian
	PixelFormatGRAY10BE PixelFormat = C.AV_PIX_FMT_GRAY10BE ///<        Y        , 10bpp, big-endian
	PixelFormatGRAY10LE PixelFormat = C.AV_PIX_FMT_GRAY10LE ///<        Y        , 10bpp, little-endian

	PixelFormatP016LE PixelFormat = C.AV_PIX_FMT_P016LE ///< like NV12, with 16bpp per component, little-endian
	PixelFormatP016BE PixelFormat = C.AV_PIX_FMT_P016BE ///< like NV12, with 16bpp per component, big-endian

	// Hardware surfaces for Direct3D11.
	//
	// This is preferred over the legacy AV_PIX_FMT_D3D11VA_VLD. The new D3D11
	// hwaccel API and filtering support AV_PIX_FMT_D3D11 only.
	//
	// data[0] contains a ID3D11Texture2D pointer, and data[1] contains the
	// texture array index of the frame as intptr_t if the ID3D11Texture2D is
	// an array texture (or always 0 if it's a normal texture).

	PixelFormatD3D11 PixelFormat = C.AV_PIX_FMT_D3D11

	PixelFormatGRAY9BE PixelFormat = C.AV_PIX_FMT_GRAY9BE ///<        Y        , 9bpp, big-endian
	PixelFormatGRAY9LE PixelFormat = C.AV_PIX_FMT_GRAY9LE ///<        Y        , 9bpp, little-endian

	PixelFormatGBRPF32BE  PixelFormat = C.AV_PIX_FMT_GBRPF32BE  ///< IEEE-754 single precision planar GBR 4:4:4,     96bpp, big-endian
	PixelFormatGBRPF32LE  PixelFormat = C.AV_PIX_FMT_GBRPF32LE  ///< IEEE-754 single precision planar GBR 4:4:4,     96bpp, little-endian
	PixelFormatGBRAPF32BE PixelFormat = C.AV_PIX_FMT_GBRAPF32BE ///< IEEE-754 single precision planar GBRA 4:4:4:4, 128bpp, big-endian
	PixelFormatGBRAPF32LE PixelFormat = C.AV_PIX_FMT_GBRAPF32LE ///< IEEE-754 single precision planar GBRA 4:4:4:4, 128bpp, little-endian

	// DRM-managed buffers exposed through PRIME buffer sharing.
	//
	// data[0] points to an AVDRMFrameDescriptor.

	PixelFormatDRM_PRIME PixelFormat = C.AV_PIX_FMT_DRM_PRIME

	// Hardware surfaces for OpenCL.
	//
	// data[i] contain 2D image objects (typed in C as cl_mem, used
	// in OpenCL as image2d_t) for each plane of the surface.

	PixelFormatOPENCL PixelFormat = C.AV_PIX_FMT_OPENCL

	PixelFormatGRAY14BE PixelFormat = C.AV_PIX_FMT_GRAY14BE ///<        Y        , 14bpp, big-endian
	PixelFormatGRAY14LE PixelFormat = C.AV_PIX_FMT_GRAY14LE ///<        Y        , 14bpp, little-endian

	PixelFormatGRAYF32BE PixelFormat = C.AV_PIX_FMT_GRAYF32BE ///< IEEE-754 single precision Y, 32bpp, big-endian
	PixelFormatGRAYF32LE PixelFormat = C.AV_PIX_FMT_GRAYF32LE ///< IEEE-754 single precision Y, 32bpp, little-endian

	PixelFormatYUVA422P12BE PixelFormat = C.AV_PIX_FMT_YUVA422P12BE ///< planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), 12b alpha, big-endian
	PixelFormatYUVA422P12LE PixelFormat = C.AV_PIX_FMT_YUVA422P12LE ///< planar YUV 4:2:2,24bpp, (1 Cr & Cb sample per 2x1 Y samples), 12b alpha, little-endian
	PixelFormatYUVA444P12BE PixelFormat = C.AV_PIX_FMT_YUVA444P12BE ///< planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), 12b alpha, big-endian
	PixelFormatYUVA444P12LE PixelFormat = C.AV_PIX_FMT_YUVA444P12LE ///< planar YUV 4:4:4,36bpp, (1 Cr & Cb sample per 1x1 Y samples), 12b alpha, little-endian

	PixelFormatNV24 PixelFormat = C.AV_PIX_FMT_NV24 ///< planar YUV 4:4:4, 24bpp, 1 plane for Y and 1 plane for the UV components, which are interleaved (first byte U and the following byte V)
	PixelFormatNV42 PixelFormat = C.AV_PIX_FMT_NV42 ///< as above, but U and V bytes are swapped

	// Vulkan hardware images.
	//
	// data[0] points to an AVVkFrame

	PixelFormatVULKAN PixelFormat = C.AV_PIX_FMT_VULKAN

	PixelFormatY210BE PixelFormat = C.AV_PIX_FMT_Y210BE ///< packed YUV 4:2:2 like YUYV422, 20bpp, data in the high bits, big-endian
	PixelFormatY210LE PixelFormat = C.AV_PIX_FMT_Y210LE ///< packed YUV 4:2:2 like YUYV422, 20bpp, data in the high bits, little-endian

	PixelFormatX2RGB10LE PixelFormat = C.AV_PIX_FMT_X2RGB10LE ///< packed RGB 10:10:10, 30bpp, (msb)2X 10R 10G 10B(lsb), little-endian, X=unused/undefined
	PixelFormatX2RGB10BE PixelFormat = C.AV_PIX_FMT_X2RGB10BE ///< packed RGB 10:10:10, 30bpp, (msb)2X 10R 10G 10B(lsb), big-endian, X=unused/undefined
	PixelFormatX2BGR10LE PixelFormat = C.AV_PIX_FMT_X2BGR10LE ///< packed BGR 10:10:10, 30bpp, (msb)2X 10B 10G 10R(lsb), little-endian, X=unused/undefined
	PixelFormatX2BGR10BE PixelFormat = C.AV_PIX_FMT_X2BGR10BE ///< packed BGR 10:10:10, 30bpp, (msb)2X 10B 10G 10R(lsb), big-endian, X=unused/undefined

	PixelFormatP210BE PixelFormat = C.AV_PIX_FMT_P210BE ///< interleaved chroma YUV 4:2:2, 20bpp, data in the high bits, big-endian
	PixelFormatP210LE PixelFormat = C.AV_PIX_FMT_P210LE ///< interleaved chroma YUV 4:2:2, 20bpp, data in the high bits, little-endian

	PixelFormatP410BE PixelFormat = C.AV_PIX_FMT_P410BE ///< interleaved chroma YUV 4:4:4, 30bpp, data in the high bits, big-endian
	PixelFormatP410LE PixelFormat = C.AV_PIX_FMT_P410LE ///< interleaved chroma YUV 4:4:4, 30bpp, data in the high bits, little-endian

	PixelFormatP216BE PixelFormat = C.AV_PIX_FMT_P216BE ///< interleaved chroma YUV 4:2:2, 32bpp, big-endian
	PixelFormatP216LE PixelFormat = C.AV_PIX_FMT_P216LE ///< interleaved chroma YUV 4:2:2, 32bpp, little-endian

	PixelFormatP416BE PixelFormat = C.AV_PIX_FMT_P416BE ///< interleaved chroma YUV 4:4:4, 48bpp, big-endian
	PixelFormatP416LE PixelFormat = C.AV_PIX_FMT_P416LE ///< interleaved chroma YUV 4:4:4, 48bpp, little-endian

	PixelFormatVUYA PixelFormat = C.AV_PIX_FMT_VUYA ///< packed VUYA 4:4:4, 32bpp, VUYAVUYA...

	PixelFormatRGBAF16BE PixelFormat = C.AV_PIX_FMT_RGBAF16BE ///< IEEE-754 half precision packed RGBA 16:16:16:16, 64bpp, RGBARGBA..., big-endian
	PixelFormatRGBAF16LE PixelFormat = C.AV_PIX_FMT_RGBAF16LE ///< IEEE-754 half precision packed RGBA 16:16:16:16, 64bpp, RGBARGBA..., little-endian

	PixelFormatVUYX PixelFormat = C.AV_PIX_FMT_VUYX ///< packed VUYX 4:4:4, 32bpp, Variant of VUYA where alpha channel is left undefined

	PixelFormatP012LE PixelFormat = C.AV_PIX_FMT_P012LE ///< like NV12, with 12bpp per component, data in the high bits, zeros in the low bits, little-endian
	PixelFormatP012BE PixelFormat = C.AV_PIX_FMT_P012BE ///< like NV12, with 12bpp per component, data in the high bits, zeros in the low bits, big-endian

	PixelFormatY212BE PixelFormat = C.AV_PIX_FMT_Y212BE ///< packed YUV 4:2:2 like YUYV422, 24bpp, data in the high bits, zeros in the low bits, big-endian
	PixelFormatY212LE PixelFormat = C.AV_PIX_FMT_Y212LE ///< packed YUV 4:2:2 like YUYV422, 24bpp, data in the high bits, zeros in the low bits, little-endian

	PixelFormatXV30BE PixelFormat = C.AV_PIX_FMT_XV30BE ///< packed XVYU 4:4:4, 32bpp, (msb)2X 10V 10Y 10U(lsb), big-endian, variant of Y410 where alpha channel is left undefined
	PixelFormatXV30LE PixelFormat = C.AV_PIX_FMT_XV30LE ///< packed XVYU 4:4:4, 32bpp, (msb)2X 10V 10Y 10U(lsb), little-endian, variant of Y410 where alpha channel is left undefined

	PixelFormatXV36BE PixelFormat = C.AV_PIX_FMT_XV36BE ///< packed XVYU 4:4:4, 48bpp, data in the high bits, zeros in the low bits, big-endian, variant of Y412 where alpha channel is left undefined
	PixelFormatXV36LE PixelFormat = C.AV_PIX_FMT_XV36LE ///< packed XVYU 4:4:4, 48bpp, data in the high bits, zeros in the low bits, little-endian, variant of Y412 where alpha channel is left undefined

	PixelFormatRGBF32BE PixelFormat = C.AV_PIX_FMT_RGBF32BE ///< IEEE-754 single precision packed RGB 32:32:32, 96bpp, RGBRGB..., big-endian
	PixelFormatRGBF32LE PixelFormat = C.AV_PIX_FMT_RGBF32LE ///< IEEE-754 single precision packed RGB 32:32:32, 96bpp, RGBRGB..., little-endian

	PixelFormatRGBAF32BE PixelFormat = C.AV_PIX_FMT_RGBAF32BE ///< IEEE-754 single precision packed RGBA 32:32:32:32, 128bpp, RGBARGBA..., big-endian
	PixelFormatRGBAF32LE PixelFormat = C.AV_PIX_FMT_RGBAF32LE ///< IEEE-754 single precision packed RGBA 32:32:32:32, 128bpp, RGBARGBA..., little-endian

	// PixelFormatP212BE PixelFormat = C.AV_PIX_FMT_P212BE ///< interleaved chroma YUV 4:2:2, 24bpp, data in the high bits, big-endian
	// PixelFormatP212LE PixelFormat = C.AV_PIX_FMT_P212LE ///< interleaved chroma YUV 4:2:2, 24bpp, data in the high bits, little-endian

	// PixelFormatP412BE PixelFormat = C.AV_PIX_FMT_P412BE ///< interleaved chroma YUV 4:4:4, 36bpp, data in the high bits, big-endian
	// PixelFormatP412LE PixelFormat = C.AV_PIX_FMT_P412LE ///< interleaved chroma YUV 4:4:4, 36bpp, data in the high bits, little-endian

	// PixelFormatGBRAP14BE PixelFormat = C.AV_PIX_FMT_GBRAP14BE ///< planar GBR 4:4:4:4 56bpp, big-endian
	// PixelFormatGBRAP14LE PixelFormat = C.AV_PIX_FMT_GBRAP14LE ///< planar GBR 4:4:4:4 56bpp, little-endian

	PixelFormatNB PixelFormat = C.AV_PIX_FMT_NB ///< number of pixel formats DO NOT USE THIS, if you want to link with shared libav* because the number of formats might differ between versions
)
