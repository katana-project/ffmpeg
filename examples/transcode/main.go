package main

import (
	"fmt"
	"github.com/katana-project/ffmpeg/avcodec"
	"github.com/katana-project/ffmpeg/avformat"
	"github.com/katana-project/ffmpeg/avutil"
	"os"
)

// based on https://github.com/FFmpeg/FFmpeg/blob/master/doc/examples/vaapi_transcode.c
// VAAPI acceleration removed

var (
	ifmtCtx     = &avformat.FormatContext{}
	ofmtCtx     = &avformat.FormatContext{}
	decoderCtx  = &avcodec.CodecContext{}
	encoderCtx  = &avcodec.CodecContext{}
	videoStream int
	ost         *avformat.Stream
	initialized bool
)

func openInputFile(filename string) (ret int) {
	if ret = ifmtCtx.OpenInput(filename, nil, nil); ret < 0 {
		fmt.Printf("Cannot open input file '%s', Error code: %s\n", filename, avutil.Err2Str(ret))
		return ret
	}

	if ret = ifmtCtx.FindStreamInfo(nil); ret < 0 {
		fmt.Printf("Cannot find input stream information. Error code: %s\n", avutil.Err2Str(ret))
		return ret
	}

	var decoder *avcodec.Codec
	ret, decoder = ifmtCtx.FindBestStream(avutil.MediaTypeVideo, -1, -1, 0)
	if ret < 0 {
		fmt.Printf("Cannot find a video stream in the input file. Error code: %s\n", avutil.Err2Str(ret))
		return ret
	}
	videoStream = ret

	if decoderCtx.AllocContext3(decoder); decoderCtx.Null() {
		return avutil.ErrorENoMem
	}

	video := ifmtCtx.Stream(videoStream)
	if ret = decoderCtx.ParametersToContext(video.CodecPar()); ret < 0 {
		fmt.Printf("avcodec_parameters_to_context error. Error code: %s\n", avutil.Err2Str(ret))
		return ret
	}

	if ret = decoderCtx.Open2(decoder, nil); ret < 0 {
		fmt.Printf("Failed to open codec for decoding. Error code: %s\n", avutil.Err2Str(ret))
	}

	return ret
}

func encodeWrite(encPkt *avcodec.Packet, frame *avutil.Frame) (ret int) {
	encPkt.Unref()
	if ret = encoderCtx.SendFrame(frame); ret < 0 {
		fmt.Printf("Error during encoding. Error code: %s\n", avutil.Err2Str(ret))

		if ret == avutil.ErrorEOF || ret == avutil.ErrorEAgain {
			return 0
		}
		return -1
	}

	for {
		if ret = encoderCtx.ReceivePacket(encPkt); ret != 0 {
			break
		}

		encPkt.SetStreamIndex(0)
		encPkt.RescaleTs(ifmtCtx.Stream(videoStream).TimeBase(), ofmtCtx.Stream(0).TimeBase())
		ret = ofmtCtx.InterleavedWriteFrame(encPkt)
		if ret < 0 {
			fmt.Printf("Error during writing data to output file. Error code: %s\n", avutil.Err2Str(ret))
			return -1
		}
	}

	if ret == avutil.ErrorEOF || ret == avutil.ErrorEAgain {
		return 0
	}
	return -1
}

func decEnc(pkt *avcodec.Packet, encCodec *avcodec.Codec) (ret int) {
	if ret = decoderCtx.SendPacket(pkt); ret < 0 {
		fmt.Printf("Error during decoding. Error code: %s\n", avutil.Err2Str(ret))
		return ret
	}

	frame := &avutil.Frame{}
	for ret >= 0 {
		if frame.Alloc(); frame.Null() {
			return avutil.ErrorENoMem
		}

		if ret = decoderCtx.ReceiveFrame(frame); ret == avutil.ErrorEAgain || ret == avutil.ErrorEOF {
			frame.Free()
			return 0
		} else if ret < 0 {
			fmt.Printf("Error while decoding. Error code: %s\n", avutil.Err2Str(ret))
			goto fail
		}

		if !initialized {
			encoderCtx.SetTimeBase(decoderCtx.Framerate().Inv())
			encoderCtx.SetPixFmt(decoderCtx.PixFmt())
			encoderCtx.SetWidth(decoderCtx.Width())
			encoderCtx.SetHeight(decoderCtx.Height())

			if ret = encoderCtx.Open2(encCodec, nil); ret < 0 {
				fmt.Printf("Failed to open encode codec. Error code: %s\n", avutil.Err2Str(ret))
				goto fail
			}

			if ost = ofmtCtx.NewStream(encCodec); ost.Null() {
				fmt.Printf("Failed to allocate stream for output format.\n")
				ret = avutil.ErrorENoMem
				goto fail
			}

			ost.SetTimeBase(encoderCtx.TimeBase())
			if ret = encoderCtx.ParametersFromContext(ost.CodecPar()); ret < 0 {
				fmt.Printf("Failed to copy the stream parameters. Error code: %s\n", avutil.Err2Str(ret))
				goto fail
			}

			if ret = ofmtCtx.WriteHeader(nil); ret < 0 {
				fmt.Printf("Error while writing stream header. Error code: %s\n", avutil.Err2Str(ret))
				goto fail
			}

			initialized = true
		}

		frame.SetPictType(avutil.PictureTypeNone)
		if ret = encodeWrite(pkt, frame); ret < 0 {
			fmt.Printf("Error during encoding and writing.\n")
		}

	fail:
		frame.Free()
		if ret < 0 {
			return ret
		}
	}
	return 0
}

func main() {
	args := os.Args
	if len(args) != 4 {
		fmt.Printf("Usage: %s <input file> <encode codec> <output file>\n"+
			"The output format is guessed according to the file extension.\n\n", args[0])
		os.Exit(-1)
	}

	var (
		encCodec *avcodec.Codec
		ret      int
		pb       *avformat.IOContext

		decPkt = &avcodec.Packet{}
	)

	if decPkt.Alloc(); decPkt.Null() {
		fmt.Println("Failed to allocate decode packet")
		goto end
	}

	if ret = openInputFile(args[1]); ret < 0 {
		goto end
	}

	if encCodec = avcodec.FindEncoderByName(args[2]); encCodec == nil {
		fmt.Printf("Could not find encoder '%s'\n", args[2])
		ret = -1
		goto end
	}

	if ret = ofmtCtx.AllocOutputContext2(nil, "", args[3]); ret < 0 {
		fmt.Printf("Failed to deduce output format from file extension. Error code: %s\n", avutil.Err2Str(ret))
		goto end
	}

	if encoderCtx.AllocContext3(encCodec); encoderCtx.Null() {
		ret = avutil.ErrorENoMem
		goto end
	}

	pb = ofmtCtx.Pb()
	if ret = pb.Open(args[3], avformat.IOFlagWrite); ret < 0 {
		fmt.Printf("Cannot open output file. Error code: %s\n", avutil.Err2Str(ret))
		goto end
	}
	ofmtCtx.SetPb(pb)

	for {
		if ret = ifmtCtx.ReadFrame(decPkt); ret < 0 {
			break
		}

		if videoStream == decPkt.StreamIndex() {
			ret = decEnc(decPkt, encCodec)
		}

		decPkt.Unref()
	}

	decPkt.Unref()
	ret = decEnc(decPkt, encCodec)

	ret = encodeWrite(decPkt, nil)

	ofmtCtx.WriteTrailer()

end:
	ifmtCtx.CloseInput()
	ofmtCtx.CloseInput()
	decoderCtx.FreeContext()
	encoderCtx.FreeContext()
	decPkt.Free()
	os.Exit(ret)
}
