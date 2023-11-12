package main

import (
	"fmt"
	"github.com/katana-project/ffmpeg/avformat"
	"github.com/katana-project/ffmpeg/avutil"
	"os"
)

// based on https://github.com/FFmpeg/FFmpeg/blob/master/doc/examples/show_metadata.c

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("usage: %s <input_file>\n", args[0])
		fmt.Println("example program to demonstrate the use of the libavformat metadata API.")
		os.Exit(1)
	}

	fmtCtx := &avformat.FormatContext{}
	if ret := fmtCtx.OpenInput(args[1], nil, nil); ret == 1 {
		os.Exit(ret)
	}

	if ret := fmtCtx.FindStreamInfo(nil); ret < 0 {
		fmt.Fprintln(os.Stderr, "Cannot find stream information")
		os.Exit(ret)
	}

	var (
		metadata = fmtCtx.Metadata()
		tag      *avutil.DictionaryEntry
	)
	for {
		tag = metadata.Iterate(tag)
		if tag == nil {
			break
		}

		fmt.Printf("%s=%s\n", tag.Key(), tag.Value())
	}

	fmtCtx.CloseInput()
}
