package main

import (
	"context"
	"demo/http"
	"demo/tracing"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	tlog "github.com/opentracing/opentracing-go/log"
)

func main() {
	tracer, closer := tracing.Init("service-1")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/call", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))

		// Span 1
		span := tracer.StartSpan("call", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		// Span 2
		span2 := tracer.StartSpan("step1", opentracing.ChildOf(span.Context()))
		span2.SetTag("step", "step 1")
		defer span2.Finish()

		// Span 3
		ctx := opentracing.ContextWithSpan(context.Background(), span2)
		callService(ctx)

		io.WriteString(w, fmt.Sprint("Call from service 1"))
	})

	http.HandleFunc("/target", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		span := tracer.StartSpan("target", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		io.WriteString(w, fmt.Sprint("Hello target"))
	})

	http.HandleFunc("/call-service2", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))

		// Span 1
		span := tracer.StartSpan("call-service2", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		// Span 2
		span2 := tracer.StartSpan("step1", opentracing.ChildOf(span.Context()))
		span2.SetTag("step", "step 1")
		defer span2.Finish()

		// Span 3
		ctx := opentracing.ContextWithSpan(context.Background(), span2)
		callService2(ctx, "hello")

		io.WriteString(w, fmt.Sprint("Demo service 1 -> 2"))
	})

	http.HandleFunc("/call-service3", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))

		// Span 1
		span := tracer.StartSpan("call-service2", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		// Span 2
		span2 := tracer.StartSpan("step1", opentracing.ChildOf(span.Context()))
		span2.SetTag("step", "step 1")
		defer span2.Finish()

		// Span 3
		ctx := opentracing.ContextWithSpan(context.Background(), span2)
		callService2(ctx, "hello3")

		io.WriteString(w, fmt.Sprint("Demo service 1 -> 2 -> 3"))
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func callService2(ctx context.Context, path string) {
	span, _ := opentracing.StartSpanFromContext(ctx, "callService")
	defer span.Finish()

	url := "http://localhost:8001/" + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	resp, err := myhttp.Do(req)
	if err != nil {
		panic(err.Error())
	}

	response := string(resp)

	span.LogFields(
		tlog.String("event", "call_service2_hello"),
		tlog.String("value", response),
	)

	fmt.Println(response)
}

func callService(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "callService")
	defer span.Finish()

	url := "http://localhost:8000/target"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err.Error())
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(req.Header),
	)

	resp, err := myhttp.Do(req)
	if err != nil {
		panic(err.Error())
	}

	response := string(resp)

	span.LogFields(
		tlog.String("event", "call_target"),
		tlog.String("value", response),
	)

	fmt.Println(response)
}
