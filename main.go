package main

import (
	"bytes"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/cheynewallace/tabby"
	"github.com/dustin/go-humanize"
)

func init() {
	if len(os.Args) != 2 {
		dieOnError("usage: %s <torrent file>", os.Args[0])
	}
}

func main() {
	var t = tabby.New()

	mi, err := metainfo.LoadFromFile(os.Args[1])
	if err != nil {
		dieOnError(err.Error())
	}

	info, err := mi.UnmarshalInfo()
	if err != nil {
		dieOnError(err.Error())
	}

	filesBuf := bytes.NewBuffer(nil)
	tf := tabby.NewCustom(tabwriter.NewWriter(filesBuf, 0, 0, 3, ' ', 0))
	for _, f := range info.UpvertedFiles() {
		tf.AddLine("    ", f.DisplayPath(&info), humanize.Bytes(uint64(f.Length)))
	}
	tf.Print()

	t.AddLine("Name", info.Name)
	t.AddLine("Size", humanize.Bytes(uint64(info.TotalLength())))
	t.AddLine("Files", "")
	t.AddLine(filesBuf)
	t.Print()
}

func dieOnError(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("error: %s\n", format), a...)
	os.Exit(1)
}
