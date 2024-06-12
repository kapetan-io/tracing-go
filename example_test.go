package tracing_test

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gubernator-io/tracing"
)

type Service struct{}

func (s *Service) callHTTP(ctx context.Context) (err error) {
	ctx = tracing.NewFuncSpan(ctx, "Service.callHTTP", err)
	defer tracing.EndSpan(ctx)

	// Pass along the context returned by `tracing.Scope()`
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://google.com", nil)
	if err != nil {
		tracing.SpanFrom(ctx).Error(err)
		return err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		tracing.SpanFrom(ctx).Error(err, "url", req.URL.String())
		return err
	}

	defer func() { _ = resp.Body.Close() }()
	b, err := io.ReadAll(resp.Body)
	fmt.Printf("Response body: %s\n", string(b))
	return nil
}

func (s *Service) Handle(w http.ResponseWriter, r *http.Request) {
	// Creates a new trace for this request, inheriting any trace request
	// id's from the upstream client if provided via headers.
	ctx := tracing.RootSpan(r.Context(), r.URL.Path, tracing.GlobalProvider,
		tracing.AttributeHttpClientIp, r.Header.Get("x-forwarded-for"),
		tracing.AttributeSpanKind, "server",
		tracing.AttributeHttpMethod, r.Method,
		tracing.AttributeHttpUrl, r.URL.String(),
	)
	defer tracing.EndSpan(ctx)

	if r.FormValue("do-event") != "" {
		tracing.SpanFrom(ctx).AddEvent("received not okay",
			"ip", r.Header.Get("Source-Ip"),
			"host", r.URL.Host,
			"method", r.Method)
	}
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
