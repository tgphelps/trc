
package trc

import (
    "fmt"
    "io"
)

const numTracePoints = 16;

var tracePt [numTracePoints]int8;

var traceNames [numTracePoints]string;

var maxTraceEntries = 0;

var trDest io.Writer = io.Discard;

func TraceOpen(dest io.Writer) {
    trDest = dest;
    tracePt[0] = 1;
    for i := 0; i < numTracePoints; i++ {
        traceNames[i] = fmt.Sprintf("TR%02d", i);
        // fmt.Printf("%s\n", traceNames[i]);
    }
}

func TraceClose() {
    trDest = io.Discard;
    tracePt[0] = 0;
}

func TraceOn(n int, name string) {
    tracePt[n] = 1;
    traceNames[n] = name;
}

func TraceOff(n int) {
    tracePt[n] = 0
}

func Tracing(n int) int8 {
    return tracePt[n];
}

func Trace(n int, format string, vals ...interface{}) {
    _, err := fmt.Fprintf(trDest, "%s ", traceNames[n])
    if err != nil {
        panic(err)
    }
    _, err = fmt.Fprintf(trDest, format, vals...);
    if err != nil {
        panic(err)
    }
    _, err = fmt.Fprintln(trDest);
    if err != nil {
        panic(err)
    }
}

func TraceIf(n int, format string, vals ...interface{}){
    if tracePt[n] == 1 {
        Trace(n, format, vals...);
    }
}

func TraceDump(n int, buff []byte) {
    if tracePt[n] == 1 {
        fmt.Fprintf(trDest, "dumping %d bytes\n", n)
    }
}

