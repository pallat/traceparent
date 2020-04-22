package traceparent

import (
	"testing"
)

func TestParse(t *testing.T) {
	want := "00-0af7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01"
	tp := Parse(want)

	get := tp.String()

	if want != get {
		t.Errorf("%s parse to %s\n", want, get)
	}
}

func TestNewSpan(t *testing.T) {
	tp := New()

	newTP := tp.NewSpan()

	if tp.sp.TraceID != newTP.sp.TraceID {
		t.Errorf("origin traceID %s <> new traceID %s\n", tp.sp.TraceIDString(), newTP.sp.TraceIDString())
	}
}
