package trc

import (
	"fmt"
	"io"

	"github.com/tgphelps/hdump"
)

const numTracePoints = 16

var tracePt [numTracePoints]bool

var traceName [numTracePoints]string

var maxTraceEntries = 0

var trDest io.Writer = io.Discard

func TraceOpen(dest io.Writer) {
	trDest = dest
	tracePt[0] = true
	for i := 0; i < numTracePoints; i++ {
		traceName[i] = fmt.Sprintf("TR%02d", i)
		// fmt.Printf("%s\n", traceName[i]);
	}
}

func TraceClose() {
	trDest = io.Discard
	tracePt[0] = false
}

func TraceOn(n int, name string) {
	tracePt[n] = true
	traceName[n] = name
}

func TraceOff(n int) {
	tracePt[n] = false
}

func Tracing(n int) bool {
	return tracePt[n]
}

func Trace(n int, format string, vals ...interface{}) {
	_, err := fmt.Fprintf(trDest, "%-6s:", traceName[n])
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintf(trDest, format, vals...)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprintln(trDest)
	if err != nil {
		panic(err)
	}
}

func TraceIf(n int, format string, vals ...interface{}) {
	if tracePt[n] {
		Trace(n, format, vals...)
	}
}

func TraceDump(n int, buff []byte) {
	if tracePt[n] == true {
		// fmt.Fprintf(trDest, "dumping %d bytes\n", n)
		dest := hdump.NewHdumper(trDest)
		dest.DumpBytes(len(buff), buff)
	}
}

func TraceInt32s(n int, buff []int32) {
	if tracePt[n] == true {
		// fmt.Fprintf(trDest, "dumping %d bytes\n", n)
		dest := hdump.NewHdumper(trDest)
		dest.DumpInt32s(len(buff), buff)
	}
}
