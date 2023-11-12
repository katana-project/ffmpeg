package main

import (
	"fmt"
	"github.com/katana-project/ffmpeg/avcodec"
	"github.com/katana-project/ffmpeg/avformat"
	"github.com/katana-project/ffmpeg/avutil"
	"os"
)

// based on https://github.com/FFmpeg/FFmpeg/blob/master/doc/examples/remux.c

func logPacket(fmtCtx *avformat.FormatContext, pkt *avcodec.Packet, tag string) {
	timeBase := fmtCtx.Stream(pkt.StreamIndex()).TimeBase()

	fmt.Printf(
		"%s: pts:%s pts_time:%s dts:%s dts_time:%s duration:%s duration_time:%s stream_index:%d\n",
		tag,
		avutil.Ts2Str(pkt.Pts()), avutil.Ts2TimeStr(pkt.Pts(), timeBase),
		avutil.Ts2Str(pkt.Dts()), avutil.Ts2TimeStr(pkt.Dts(), timeBase),
		avutil.Ts2Str(pkt.Duration()), avutil.Ts2TimeStr(pkt.Duration(), timeBase),
		pkt.StreamIndex(),
	)
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Printf("usage: %s input output\n", args[0])
		fmt.Println("API example program to remux a media file with libavformat and libavcodec.")
		fmt.Println("The output format is guessed according to the file extension.")
		os.Exit(1)
	}

	var (
		ifmtCtx = &avformat.FormatContext{}
		ofmtCtx = &avformat.FormatContext{}

		ofmt *avformat.OutputFormat

		inFilename  = args[1]
		outFilename = args[2]
	)

	pkt := &avcodec.Packet{}
	if pkt.Null() {
		fmt.Fprintln(os.Stderr, "Could not allocate AVPacket")
		os.Exit(1)
	}

	endFunc := func(ret int) {
		pkt.Free()
		ifmtCtx.CloseInput()

		if !ofmtCtx.Null() && (ofmt.Flags()&avformat.FmtNoFile) == 0 {
			pb := ofmtCtx.Pb()
			pb.CloseP()

			ofmtCtx.SetPb(pb)
		}
		ofmtCtx.FreeContext()

		if ret < 0 && ret != avutil.ErrorEOF {
			fmt.Fprintf(os.Stderr, "Error occurred: %s\n", avutil.Err2Str(ret))
			os.Exit(1)
		}

		os.Exit(0)
	}

	if ret := ifmtCtx.OpenInput(inFilename, nil, nil); ret < 0 {
		fmt.Fprintf(os.Stderr, "Could not open input file '%s'\n", inFilename)
		endFunc(ret)
	}

	if ret := ifmtCtx.FindStreamInfo(nil); ret < 0 {
		fmt.Fprintln(os.Stderr, "Failed to retrieve input stream information")
		endFunc(ret)
	}

	ifmtCtx.DumpFormat(0, inFilename, false)

	ofmtCtx.AllocOutputContext2(nil, "", outFilename)
	if ofmtCtx.Null() {
		fmt.Fprintln(os.Stderr, "Could not create output context")
		endFunc(avutil.ErrorUnknown)
	}

	ofmt = ofmtCtx.OFormat()
	var (
		inStreams     = ifmtCtx.Streams()
		streamIndex   = 0
		streamMapping = make([]int, ifmtCtx.NbStreams())
	)
	for i, inStream := range inStreams {
		inCodecPar := inStream.CodecPar()

		codecType := inCodecPar.CodecType()
		if codecType != avutil.MediaTypeAudio && codecType != avutil.MediaTypeVideo && codecType != avutil.MediaTypeSubtitle {
			streamMapping[i] = -1
			continue
		}

		streamMapping[i] = streamIndex
		streamIndex++

		outStream := ofmtCtx.NewStream(nil)
		if outStream.Null() {
			fmt.Fprintln(os.Stderr, "Failed allocating output stream")
			endFunc(avutil.ErrorUnknown)
		}

		outCodecPar := outStream.CodecPar()
		if ret := inCodecPar.Copy(outCodecPar); ret < 0 {
			fmt.Fprintln(os.Stderr, "Failed to copy codec parameters")
			endFunc(ret)
		}
		outCodecPar.SetCodecTag(0)
	}
	ofmtCtx.DumpFormat(0, outFilename, true)

	if (ofmt.Flags() & avformat.FmtNoFile) == 0 {
		pb := ofmtCtx.Pb()
		if ret := pb.Open(outFilename, avformat.IOFlagWrite); ret < 0 {
			fmt.Fprintf(os.Stderr, "Could not open output file '%s'\n", outFilename)
			endFunc(ret)
		}

		ofmtCtx.SetPb(pb)
	}

	if ret := ofmtCtx.WriteHeader(nil); ret < 0 {
		fmt.Fprintln(os.Stderr, "Error occurred when opening output file")
		endFunc(ret)
	}

	outStreams := ofmtCtx.Streams()
	for {
		if ret := ifmtCtx.ReadFrame(pkt); ret < 0 {
			break
		}

		var (
			streamIdx = pkt.StreamIndex()
			inStream  = inStreams[streamIdx]
			mappedId  = streamMapping[streamIdx]
		)
		if streamIdx >= streamIndex || mappedId < 0 {
			pkt.Unref()
			continue
		}

		pkt.SetStreamIndex(mappedId)
		outStream := outStreams[mappedId]
		logPacket(ifmtCtx, pkt, "in")

		pkt.RescaleTs(inStream.TimeBase(), outStream.TimeBase())
		pkt.SetPos(-1)
		logPacket(ofmtCtx, pkt, "out")

		if ret := ofmtCtx.InterleavedWriteFrame(pkt); ret < 0 {
			fmt.Fprintln(os.Stderr, "Error muxing packet")
			break
		}
	}

	ofmtCtx.WriteTrailer()
	endFunc(0)
}
