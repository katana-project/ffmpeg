package main

import (
	"fmt"
	"github.com/katana-project/ffmpeg"
	"github.com/katana-project/ffmpeg/avformat"
)

func main() {
	muxIter := &ffmpeg.IterationState{}
	for {
		of := avformat.MuxerIterate(muxIter)
		if of == nil {
			break
		}

		fmt.Printf(
			"mux flags:%d name:%s long_name:%s mime_type:%s extensions:%s\n",
			of.Flags(), of.Name(), of.LongName(), of.MIMEType(), of.Extensions(),
		)
	}

	demuxIter := &ffmpeg.IterationState{}
	for {
		ifo := avformat.DemuxerIterate(demuxIter)
		if ifo == nil {
			break
		}

		fmt.Printf(
			"demux flags:%d name:%s long_name:%s mime_type:%s extensions:%s\n",
			ifo.Flags(), ifo.Name(), ifo.LongName(), ifo.MIMEType(), ifo.Extensions(),
		)
	}
}
