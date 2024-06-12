package tracing

import (
	"context"
)

type noopProvider struct{}

type onErrorProvider struct{}

func (o onErrorProvider) NewSpan(parent Span, name string, attrs ...any) Span {
	//TODO implement me
	return nil
}

type OnErrorOptions struct{}

// NewOnErrorProvider reports spans only if sampled or if an error is present in the span.
func NewOnErrorProvider(ctx context.Context, opts OnErrorOptions) Provider {
	// TODO: Configure the sampler
	// TODO: Record all the Provider level attributes

	return &onErrorProvider{}
}

// TODO(thrawn01): I don't like the name Provider as that isn't really what this does.

// Provider decides if this trace should be sampled or not, and sets has all the basic attributes which
// should be associated with all spans
type Provider interface {
	NewSpan(parent Span, name string, attrs ...any) Span
}

var GlobalProvider = &onErrorProvider{}
