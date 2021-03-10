package main

import (
	"demo/tracing"
	"demo/http"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	tlog "github.com/opentracing/opentracing-go/log"
)

func main() {
	tracer, closer := tracing.Init("service-2")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))

		// Span 1
		span := tracer.StartSpan("hello", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		// Span 2
		span2 := tracer.StartSpan("step1-hello", opentracing.ChildOf(span.Context()))
		span2.SetTag("step", "step 1")
		defer span2.Finish()

		// Span 3
		span3 := tracer.StartSpan("step2-hello", opentracing.ChildOf(span2.Context()))
		span3.SetTag("step", "step 2")
		defer span3.Finish()

		io.WriteString(w, fmt.Sprint("Hello from service 2"))
	})

	http.HandleFunc("/hello3", func(w http.ResponseWriter, r *http.Request) {
		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))

		// Span 1
		span := tracer.StartSpan("hello", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		// Span 2
		span2 := tracer.StartSpan("step1-hello", opentracing.ChildOf(span.Context()))
		span2.SetTag("step", "step 1")
		defer span2.Finish()

		// Span 3
		ctx := opentracing.ContextWithSpan(context.Background(), span2)
		callService3(ctx)

		io.WriteString(w, fmt.Sprint("Hello from service 2"))
	})

	log.Fatal(http.ListenAndServe(":8001", nil))
}

func callService3(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "callService")
	defer span.Finish()

	url := "http://localhost:8002/call"
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
		tlog.String("event", "call_service3"),
		tlog.String("value", response),
	)

	fmt.Println(response)
}
