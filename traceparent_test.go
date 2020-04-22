package traceparent

import (
	"fmt"
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

	t.Run("", func(t *testing.T) {
		if tp.TraceID != newTP.TraceID {
			t.Errorf("origin traceID %s <> new traceID %s\n", tp.TraceIDString(), newTP.TraceIDString())
		}
	})
	t.Run("", func(t *testing.T) {
		if tp.SpanIDString() == newTP.SpanIDString() {
			t.Errorf("origin SpanID %s == new SpanID %s\n", tp.TraceIDString(), newTP.TraceIDString())
		}
	})
}

func TestX(t *testing.T) {
	tp := Parse("00-0af7651916cd43dd8448eb211c80319c-b7ad6b7169203331-01")

	span := tp.NewSpan()

	fmt.Println(tp, span)
}
