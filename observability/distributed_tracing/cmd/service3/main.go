package main

import (
	"demo/tracing"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func main() {
	tracer, closer := tracing.Init("service-3")
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

		io.WriteString(w, fmt.Sprint("Hello service 3"))
	})

	log.Fatal(http.ListenAndServe(":8002", nil))
}

