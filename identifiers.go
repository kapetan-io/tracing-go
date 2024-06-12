// Copyright The OpenTelemetry Authors
// TODO(thrawn01): License

package tracing

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
)

const (
	errInvalidHexID errorConst = "trace-id and span-id can only contain [0-9a-f] characters, all lowercase"

	errInvalidTraceIDLength errorConst = "hex encoded trace-id must have length equals to 32"
	errNilTraceID           errorConst = "trace-id can't be all zero"

	errInvalidSpanIDLength errorConst = "hex encoded span-id must have length equals to 16"
	errNilSpanID           errorConst = "span-id can't be all zero"
)

type errorConst string

func (e errorConst) Error() string {
	return string(e)
}

// TraceID is a unique identity of a trace.
// nolint:revive // revive complains about stutter of `trace.TraceID`.
type TraceID [16]byte

var (
	nilTraceID TraceID
	_          json.Marshaler = nilTraceID
)

// IsValid checks whether the trace TraceID is valid. A valid trace ID does
// not consist of zeros only.
func (t TraceID) IsValid() bool {
	return !bytes.Equal(t[:], nilTraceID[:])
}

// MarshalJSON implements a custom marshal function to encode TraceID
// as a hex string.
func (t TraceID) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// String returns the hex string representation form of a TraceID.
func (t TraceID) String() string {
	return hex.EncodeToString(t[:])
}

// SpanID is a unique identity of a span in a trace.
type SpanID [8]byte

var (
	nilSpanID SpanID
	_         json.Marshaler = nilSpanID
)

// IsValid checks whether the SpanID is valid. A valid SpanID does not consist
// of zeros only.
func (s SpanID) IsValid() bool {
	return !bytes.Equal(s[:], nilSpanID[:])
}

// MarshalJSON implements a custom marshal function to encode SpanID
// as a hex string.
func (s SpanID) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// String returns the hex string representation form of a SpanID.
func (s SpanID) String() string {
	return hex.EncodeToString(s[:])
}

// TraceIDFromHex returns a TraceID from a hex string if it is compliant with
// the W3C trace-context specification.  See more at
// https://www.w3.org/TR/trace-context/#trace-id
// nolint:revive // revive complains about stutter of `trace.TraceIDFromHex`.
func TraceIDFromHex(h string) (TraceID, error) {
	t := TraceID{}
	if len(h) != 32 {
		return t, errInvalidTraceIDLength
	}

	if err := decodeHex(h, t[:]); err != nil {
		return t, err
	}

	if !t.IsValid() {
		return t, errNilTraceID
	}
	return t, nil
}

// SpanIDFromHex returns a SpanID from a hex string if it is compliant
// with the w3c trace-context specification.
// See more at https://www.w3.org/TR/trace-context/#parent-id
func SpanIDFromHex(h string) (SpanID, error) {
	s := SpanID{}
	if len(h) != 16 {
		return s, errInvalidSpanIDLength
	}

	if err := decodeHex(h, s[:]); err != nil {
		return s, err
	}

	if !s.IsValid() {
		return s, errNilSpanID
	}
	return s, nil
}

func decodeHex(h string, b []byte) error {
	for _, r := range h {
		switch {
		case 'a' <= r && r <= 'f':
			continue
		case '0' <= r && r <= '9':
			continue
		default:
			return errInvalidHexID
		}
	}

	decoded, err := hex.DecodeString(h)
	if err != nil {
		return err
	}

	copy(b, decoded)
	return nil
}
