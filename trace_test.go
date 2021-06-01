
package trc

import (
    "strings"
    "testing"
    "os"
    // "fmt"
)

var buffer  [8]byte;
var b strings.Builder;

func TestExist(t *testing.T) {
    TraceOpen(os.Stdout);
    TraceClose();
    TraceOn(0, "ALWAYS");
    TraceOff(0);
    _ = Tracing(0);
    Trace(0, "abc");
    TraceIf(0, "abc");
    TraceDump(0, buffer[0:8]);
}

func Test1(t *testing.T) {
    TraceOpen(&b);
    TraceOn(0, "TEST1");
    TraceIf(0, "hello world");
    TraceIf(1, "error world");
    TraceOff(0)
    TraceIf(0, "error world");
    TraceClose();
    // fmt.Println(b.String());
    if b.String() != "TEST1 hello world\n" {
        t.Errorf("Test1: %s", b.String());
    }
}

