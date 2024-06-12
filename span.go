package tracing

import (
	"context"
	"net/http"
)

// Span is the individual component of a trace. It represents a single named
// and timed operation of a workflow that is traced. A Tracer is used to
// create a Span and it is then up to the operation the Span represents to
// properly end the Span when the operation itself ends.
//
// Warning: Methods may be added to this interface in minor releases. See
// package documentation on API implementation for information on how to set
// default behavior for unimplemented methods.
type Span interface {
	// AddEvent adds an event with the provided name and optional attributes.
	AddEvent(name string, attrs ...any)

	// AddLink adds a link.
	// TODO(thrawn01): This might be needed, so I should keep this in mind as a future addition.
	// AddLink(link Link)

	// IsSampling returns true if the span is being sampled.
	// true if the Span is active and events can be recorded.
	IsSampling() bool

	// Error will record err as an exception span event for this span and
	// attach any additional attributes provided. It will also set the attribute
	// AttributeHasError to true
	Error(err error, attrs ...any)

	// SetName sets the Span name.
	SetName(name string)

	// SetAttributes sets attributes of the Span. If an attribute already
	// exists it will be overwritten with the value provided.
	SetAttributes(attrs ...any)

	// Provider returns a Provider that can be used to generate
	// additional Spans on the same telemetry pipeline as the current Span.
	// TODO(thrawn01): This might not be needed
	Provider() Provider

	// TraceID returns the trace id of this span, the id of the root span this span is associated with.
	TraceID() TraceID

	// SpanID returns the span id unique to this span.
	SpanID() SpanID
}

// SpanFrom creates a new span using the provided context
func SpanFrom(ctx context.Context) Span {
	return nil
}

// RootSpan creates a new root span which will report to the trace provider once the root span is finalized.
// A trace is just a bunch of spans that all relate to each other. This tracer holds on to those spans until
// it needs to send them to the trace provider for processing.
func RootSpan(ctx context.Context, name string, provider Provider, attrs ...any) context.Context {
	return nil
}

// RootSpanFromRequest creates a new root span like RootSpan(). In addition, it extracts trace information
// from the request headers and adds common span attributes like `http.url` to the root span.
func RootSpanFromRequest(r *http.Request, provider Provider, attrs ...any) context.Context {
	return nil
}

// NewFuncSpan is a convenience function which is intended to be called at the top of a function and
// reports if the function returned an error.
func NewFuncSpan(ctx context.Context, name string, err error) context.Context {
	return ctx
}

func NewSpan(ctx context.Context, name string, attrs ...any) context.Context {
	return ctx
}

func EndSpan(ctx context.Context, attrs ...any) {

}
