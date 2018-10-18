// +build !noruntimetrace

package tracing

import (
	"context"
	"runtime/trace"

	opentracing "github.com/opentracing/opentracing-go"
)

type Span struct {
	span      opentracing.Span
	region    *trace.Region
	task      *trace.Task
	finishers []func()
}

// parent span => trace.NewTask()
// inner span => trace.NewRegion()
// set tag => trace.log
type activeSpanKey struct{}

// NewSpan creates a new tracing span with no link to an existing opentracing Span.
func NewSpan(ctx context.Context, name string) (*Span, context.Context) {
	ctx, task := trace.NewTask(ctx, name)
	sp := &Span{
		task:      task,
		finishers: []func(){task.End},
	}
	return sp, ctx
}

// StartSpanFromContext
func StartSpanFromContext(ctx context.Context, operationName string) (*Span, context.Context) {
	var sp *Span

	otSpan, ctx := opentracing.StartSpanFromContext(ctx, operationName)
	if sp = SpanFromContext(ctx); sp == nil {
		sp, ctx = NewSpan(ctx, operationName)
	} else {
		// A Span already exists so just create a new region.
		region := trace.StartRegion(ctx, operationName)
		sp = &Span{region: region, finishers: []func(){region.End}}
	}
	sp.span = otSpan
	sp.finishers = append(sp.finishers, otSpan.Finish)
	return sp, ctx
}

// SpanFromContext returns the current Span from the context.
func SpanFromContext(ctx context.Context) *Span {
	val := ctx.Value(activeSpanKey{})
	if sp, ok := val.(*Span); ok {
		return sp
	}
	return nil
}

func (s *Span) SetTag(ctx context.Context, key string, value interface{}) *Span {
	if str, ok := value.(string); ok {
		trace.Log(ctx, key, str)
	}
	if s.span != nil {
		s.span.SetTag(key, value)
	}
	return s
}

func (s *Span) Finish() {
	for _, f := range s.finishers {
		f()
	}
}
