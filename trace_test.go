package trc

import (
	// "fmt"

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
	TraceOn(0, "TST1")
	TraceIf(0, "hello world")
	TraceIf(1, "error world")
	TraceOff(0)
	TraceIf(0, "error world")
	TraceClose()
	// fmt.Println(b.String());
	if b.String() != "TST1  :hello world\n" {
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
	TraceOn(0, "DUMP")
	TraceIf(0, "hex dump")
	TraceDump(0, buffer[0:8])
	TraceClose()
	// fmt.Println(b.String())
	if b.String() != "DUMP  :hex dump\n00000000 61 62 63 64 65 66 67 68                           abcdefgh\n" {
		t.Errorf("TestDump: %s", b.String())
	}
}

func TestInts(t *testing.T) {
	var results = "INTS  :hex dump\n" +
		"00000000 00000001 00000002 00000003 00000004 00000005 00000006 00000007 00000008 \n" +
		"00000008 000000ff 0000ffff 00ffffff ffffffff \n"
	var b strings.Builder
	buffer := []int32{1, 2, 3, 4, 5, 6, 7, 8, 255, 65535, (1 << 24) - 1, -1}
	TraceOpen(&b)
	TraceOn(0, "INTS")
	TraceIf(0, "hex dump")
	TraceInt32s(0, buffer[:])
	TraceClose()
	//fmt.Println(b.String())
	if b.String() != results {
		t.Errorf("TestDump: %s", b.String())
	}
}
