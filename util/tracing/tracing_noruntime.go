// +build noruntimetrace

package tracing

import (
	"context"

	opentracing "github.com/opentracing/opentracing-go"
)

type Span struct{ otSpan opentracing.Span }

// NewSpan creates a new tracing span with no link to an existing opentracing Span.
func NewSpan(ctx context.Context, name string) (*Span, context.Context) {
	return &Span{}, ctx
}

// StartSpanFromContext
func StartSpanFromContext(ctx context.Context, operationName string) (*Span, context.Context) {
	otSpan, ctx := opentracing.StartSpanFromContext(ctx, operationName)
	return &Span{otSpan: otSpan}, ctx
}

func SpanFromContext(ctx context.Context) *Span {
	return &Span{otSpan: opentracing.SpanFromContext(ctx)}
}

func (s *Span) SetTag(ctx context.Context, key string, value interface{}) *Span {
	if s.otSpan != nil {
		s.otSpan.SetTag(key, value)
	}
	return s
}

func (s *Span) Finish() {
	if s.otSpan != nil {
		s.otSpan.Finish()
	}
}
