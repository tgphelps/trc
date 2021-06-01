package trc

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var buffer [8]byte
var b strings.Builder

func TestExist(t *testing.T) {
	TraceOpen(os.Stdout)
	TraceClose()
	TraceOn(0, "ALWAYS")
	TraceOff(0)
	_ = Tracing(0)
	Trace(0, "abc")
	TraceIf(0, "abc")
	TraceDump(0, buffer[0:8])
}

func Test1(t *testing.T) {
	TraceOpen(&b)
	TraceOn(0, "TEST1")
	TraceIf(0, "hello world")
	TraceIf(1, "error world")
	TraceOff(0)
	TraceIf(0, "error world")
	TraceClose()
	// fmt.Println(b.String());
	if b.String() != "TEST1 hello world\n" {
		t.Errorf("Test1: %s", b.String())
	}
}

func Test2(t *testing.T) {
	var b strings.Builder
	TraceOpen(&b)
	TraceClose()
	TraceOn(0, "TEST1")
	TraceIf(0, "error world")
	// fmt.Println(b.String());
	if b.String() != "" {
		t.Errorf("Test2: %s", b.String())
	}
}

func TestDump(t *testing.T) {
	var b strings.Builder
	buffer := []byte("abcdefgh")
	TraceOpen(&b)
	TraceOn(0, "DUMPER")
	TraceIf(0, "hex dump")
	TraceDump(0, buffer[0:8])
	TraceClose()
	fmt.Println(b.String())
	if b.String() != "DUMPER hex dump\n00000000 61 62 63 64 65 66 67 68                           abcdefgh\n" {
		t.Errorf("TestDump: %s", b.String())
	}
}
