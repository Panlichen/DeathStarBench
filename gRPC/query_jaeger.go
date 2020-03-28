package main

import (
	"context"
	"fmt"
	"io"
	"log"

	jaeger_pb "github.com/jaegertracing/jaeger/proto-gen/api_v2"
	"google.golang.org/grpc"
)

func main() {
	log.Println("begin")
	// conn, err := grpc.Dial("10.68.229.36:80", grpc.WithInsecure())
	conn, err := grpc.Dial("10.68.241.59:16686", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	log.Println("connected")
	qsClient := jaeger_pb.NewQueryServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// cannot determine TraceID
	// var in = &jaeger_pb.GetTraceRequest{TraceID: jaeger_model.NewTraceID(1, 2)}

	// qsgtClient, err2 := qsClient.GetTrace(ctx, in)

	ftRequest := &jaeger_pb.FindTracesRequest{
		Query: &jaeger_pb.TraceQueryParameters{
			ServiceName:   "sn-nginx",
			OperationName: "Follow",
			Tags:          map[string]string{"internal.span.format": "proto"},
		},
	}
	qsftClient, err2 := qsClient.FindTraces(ctx, ftRequest)

	if err2 != nil {
		log.Fatalf("GetTrace failed: %v", err2)
	}
	log.Println("get stream")
	for {
		spanResposeChunk, err3 := qsftClient.Recv()
		if err3 != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err3)
		}
		spans := spanResposeChunk.GetSpans()
		fmt.Println(len(spans))
		// for i, span := range spans {
		// 	fmt.Printf("== Span #%v==\n\tTraceID: hi: %v, lo: %v; SpanID: %v, OperationName: %v, "+
		// 		"Flags: %v; StartTime: %v; \n",
		// 		i, span.TraceID.High, span.TraceID.Low, span.SpanID, span.OperationName, span.Flags,
		// 		span.StartTime)
		// }
	}
}
