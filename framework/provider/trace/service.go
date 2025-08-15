package trace

import (
	"context"
	"github.com/orgming/ming/framework"
	"github.com/orgming/ming/framework/contract"
	"github.com/orgming/ming/framework/gin"
	"net/http"
	"time"
)

type TraceKey string

var ContextKey = TraceKey("trace-key")

type MingTraceService struct {
	idService contract.IDService

	traceIDGenerator contract.IDService
	spanIDGenerator  contract.IDService
}

func NewMingTraceService(params ...any) (any, error) {
	c := params[0].(framework.Container)
	idService := c.MustMake(contract.IDKey).(contract.IDService)
	return &MingTraceService{idService: idService}, nil
}

func (t *MingTraceService) WithTrace(c context.Context, trace *contract.TraceContext) context.Context {
	if ginC, ok := c.(*gin.Context); ok {
		ginC.Set(string(ContextKey), trace)
		return ginC
	} else {
		newC := context.WithValue(c, ContextKey, trace)
		return newC
	}
}

func (t *MingTraceService) GetTrace(c context.Context) *contract.TraceContext {
	if ginC, ok := c.(*gin.Context); ok {
		if val, ok2 := ginC.Get(string(ContextKey)); ok2 {
			return val.(*contract.TraceContext)
		}
	}

	if tc, ok := c.Value(ContextKey).(*contract.TraceContext); ok {
		return tc
	}
	return nil
}

func (t *MingTraceService) NewTrace() *contract.TraceContext {
	var traceID, spanID string
	if t.traceIDGenerator != nil {
		traceID = t.traceIDGenerator.NewID()
	} else {
		traceID = t.idService.NewID()
	}

	if t.spanIDGenerator != nil {
		spanID = t.spanIDGenerator.NewID()
	} else {
		spanID = t.idService.NewID()
	}
	tc := &contract.TraceContext{
		TraceID:    traceID,
		ParentID:   "",
		SpanID:     spanID,
		CspanID:    "",
		Annotation: map[string]string{},
	}
	return tc
}

func (t *MingTraceService) StartSpan(trace *contract.TraceContext) *contract.TraceContext {
	var childSpanID string
	if t.spanIDGenerator != nil {
		childSpanID = t.spanIDGenerator.NewID()
	} else {
		childSpanID = t.idService.NewID()
	}
	childSpan := &contract.TraceContext{
		TraceID:  trace.TraceID,
		ParentID: "",
		SpanID:   trace.SpanID,
		CspanID:  childSpanID,
		Annotation: map[string]string{
			contract.TraceKeyTime: time.Now().String(),
		},
	}
	return childSpan
}

func (t *MingTraceService) ToMap(trace *contract.TraceContext) map[string]string {
	m := map[string]string{}
	if trace == nil {
		return m
	}
	m[contract.TraceKeyTraceID] = trace.TraceID
	m[contract.TraceKeySpanID] = trace.SpanID
	m[contract.TraceKeyCspanID] = trace.CspanID
	m[contract.TraceKeyParentID] = trace.ParentID

	if trace.Annotation != nil {
		for k, v := range trace.Annotation {
			m[k] = v
		}
	}
	return m
}

func (t *MingTraceService) ExtractHTTP(r *http.Request) *contract.TraceContext {
	trace := &contract.TraceContext{}
	trace.TraceID = r.Header.Get(contract.TraceKeyTraceID)
	trace.ParentID = r.Header.Get(contract.TraceKeyParentID)
	trace.SpanID = r.Header.Get(contract.TraceKeySpanID)
	trace.CspanID = r.Header.Get(contract.TraceKeyCspanID)

	if trace.TraceID == "" {
		trace.TraceID = t.idService.NewID()
	}

	if trace.SpanID == "" {
		trace.SpanID = t.idService.NewID()
	}

	return trace
}

func (t *MingTraceService) InjectHTTP(r *http.Request, trace *contract.TraceContext) *http.Request {
	r.Header.Add(contract.TraceKeyTraceID, trace.TraceID)
	r.Header.Add(contract.TraceKeyParentID, trace.ParentID)
	r.Header.Add(contract.TraceKeySpanID, trace.SpanID)
	r.Header.Add(contract.TraceKeyCspanID, trace.CspanID)
	return r
}
