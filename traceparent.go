package traceparent

import (
	crand "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"sync"

	"go.opentelemetry.io/otel/api/core"
)

const supportedVersion = 0

// TraceParent composite go.opentelemetry.io/otel/api/core SpanContext
type TraceParent struct{ core.SpanContext }

// New returns *TraceParent
func New() TraceParent {
	sp := core.EmptySpanContext()
	id := defIDGenerator()
	sp.TraceID = id.NewTraceID()
	sp.SpanID = id.NewSpanID()

	return TraceParent{sp}
}

// Parse parse traceparent string to TraceParent
func Parse(parent string) TraceParent {
	if parent == "" {
		return New()
	}

	token := strings.Split(parent, "-")
	sp := core.EmptySpanContext()

	traceIDBytes, _ := hex.DecodeString(token[1])
	copy(sp.TraceID[:], traceIDBytes[:16])

	spanIDBytes, _ := hex.DecodeString(token[2])
	copy(sp.SpanID[:], spanIDBytes[:8])

	traceFlagsBytes, _ := hex.DecodeString(token[3])
	sp.TraceFlags = traceFlagsBytes[0]

	return TraceParent{sp}
}

func (tp TraceParent) String() string {
	return fmt.Sprintf("%.2x-%s-%.16x-%.2x",
		supportedVersion,
		tp.TraceIDString(),
		tp.SpanID,
		tp.TraceFlags&core.TraceFlagsSampled)
}

// NewSpan make new spanID with same traceID
func (tp *TraceParent) NewSpan() TraceParent {
	sp := core.EmptySpanContext()
	sp.TraceID = tp.TraceID

	id := defIDGenerator()
	sp.SpanID = id.NewSpanID()

	sp.TraceFlags = tp.TraceFlags

	return TraceParent{sp}
}

func defIDGenerator() *defaultIDGenerator {
	gen := &defaultIDGenerator{}
	var rngSeed int64
	_ = binary.Read(crand.Reader, binary.LittleEndian, &rngSeed)
	gen.randSource = rand.New(rand.NewSource(rngSeed))
	return gen
}

type defaultIDGenerator struct {
	sync.Mutex
	randSource *rand.Rand
}

func (gen *defaultIDGenerator) NewSpanID() core.SpanID {
	gen.Lock()
	defer gen.Unlock()
	sid := core.SpanID{}
	gen.randSource.Read(sid[:])
	return sid
}

// NewTraceID returns a non-zero trace ID from a randomly-chosen sequence.
// mu should be held while this function is called.
func (gen *defaultIDGenerator) NewTraceID() core.TraceID {
	gen.Lock()
	defer gen.Unlock()
	tid := core.TraceID{}
	gen.randSource.Read(tid[:])
	return tid
}
