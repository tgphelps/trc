
package trc

import (
    "strings"
    "testing"
)

var buffer  [8]byte;

func TestTrace(t *testing.T) {
    var b strings.Builder;
    TraceOpen(&b);
    TraceClose();
    TraceOn(0, "ALWAYS");
    TraceOff(0);
    _ = Tracing(0);
    Trace(0, "abc");
    TraceIf(0, "abc");
    TraceDump(0, buffer[0:8]);
}
