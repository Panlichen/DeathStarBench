package main

import (
	"context"
	"fmt"
	"io"
	"log"

	jaeger_pb "github.com/jaegertracing/jaeger/proto-gen/api_v2"
	"google.golang.org/grpc"
)

func findTraces(ctx context.Context, qsClient jaeger_pb.QueryServiceClient) {
	// ftRequest := &jaeger_pb.FindTracesRequest{Query: &jaeger_pb.TraceQueryParameters{}}
	ftRequest := &jaeger_pb.FindTracesRequest{
		Query: &jaeger_pb.TraceQueryParameters{
			ServiceName: "sn-nginx",
			// OperationName: "Follow",
			// Tags:          map[string]string{"internal.span.format": "proto"},
		},
	}
	qsftClient, err := qsClient.FindTraces(ctx, ftRequest)

	if err != nil {
		log.Fatalf("GetTrace failed: %v", err)
	}
	log.Println("get stream")
	totalSpanCnt := 0
	// for {
	spanResposeChunk, err := qsftClient.Recv()
	if err != nil {
		if err == io.EOF {
			// break
		}
		log.Fatal(err)
	}
	spans := spanResposeChunk.GetSpans()
	totalSpanCnt += len(spans)
	for i, span := range spans {
		fmt.Printf("== Span #%v==\n\tTraceID: (%v, %v); SpanID: %v; Process.ServiceName: %v; "+
			"OperationName: %v; StartTime: %v; Duration: %v;\n",
			i, span.TraceID.High, span.TraceID.Low, span.SpanID, span.Process.GetServiceName(),
			span.OperationName, span.StartTime.UnixNano()/1e3, span.Duration.Microseconds())
		for j, spanRef := range span.GetReferences() {
			fmt.Printf("\t=SpanRef #%v= TraceID: (%v, %v); SpanID: %v; SpanRefType: %v\n",
				j, spanRef.TraceID.High, spanRef.TraceID.Low, spanRef.SpanID, spanRef.RefType)
		}
		for j, pTag := range span.Process.GetTags() {
			fmt.Printf("\t~pTag #%v~ %v = %v;\n",
				j, pTag.Key, pTag.VStr)
		}
		println()
	}
	// }
	fmt.Printf("Total Span #: %v\n", totalSpanCnt)
}

func main() {
	log.Println("begin")
	// conn, err := grpc.Dial("10.68.229.36:80", grpc.WithInsecure())
	conn, err := grpc.Dial("10.0.0.252:31010", grpc.WithInsecure())
	// conn, err := grpc.Dial("172.20.232.17:16686", grpc.WithInsecure())
	// conn, err := grpc.Dial("10.68.241.59:16686", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	log.Println("connected")
	qsClient := jaeger_pb.NewQueryServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	findTraces(ctx, qsClient)
}
