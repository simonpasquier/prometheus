package tracing

import (
	"context"
	"runtime/trace"

	opentracing "github.com/opentracing/opentracing-go"
)

type Span struct {
	spans   []opentracing.Span
	regions []*trace.Region
	task    *trace.Task
}

// parent span => trace.NewTask()
// inner span => trace.NewRegion()
// set tag => trace.log
type activeSpanKey struct{}

// NewSpan creates a new tracing span with no link to an existing opentracing Span.
func NewSpan(ctx context.Context, name string) (*Span, context.Context) {
	ctx, task := trace.NewTask(ctx, name)
	sp := &Span{
		task:    task,
		spans:   make([]opentracing.Span, 0),
		regions: make([]*trace.Region, 0),
	}
	return sp, ctx
}

// StartSpanFromContext
func StartSpanFromContext(ctx context.Context, operationName string) (*Span, context.Context) {
	var (
		otSpan opentracing.Span
		sp     *Span
	)

	otSpan, ctx = opentracing.StartSpanFromContext(ctx, operationName)
	if sp = SpanFromContext(ctx); sp == nil {
		sp, ctx = NewSpan(ctx, operationName)
		sp.spans = append(sp.spans, otSpan)
	} else {
		// A Span already exists so just create a new region.
		region := trace.StartRegion(ctx, operationName)
		sp.regions = append(sp.regions, region)
	}
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
	if len(s.spans) > 0 {
		s.spans[len(s.spans)-1].SetTag(key, value)
	}
	return s
}

func (s *Span) Finish() {
	if len(s.spans) > 0 {
		s.spans[len(s.spans)-1].Finish()
		s.spans = s.spans[:len(s.spans)-1]
	}
	if len(s.regions) > 0 {
		s.regions[len(s.regions)-1].End()
		s.regions = s.regions[:len(s.regions)-1]
	} else {
		s.task.End()
	}
}
