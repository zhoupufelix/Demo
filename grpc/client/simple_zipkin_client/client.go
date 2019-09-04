package main

import (
		"context"
		"google.golang.org/grpc"
		"log"
		pb "Demo/grpc/proto"
		"google.golang.org/grpc/credentials"
		"os"
		zipkin "github.com/openzipkin/zipkin-go-opentracing"
		"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
)

const (
	PORT = "9007"
	SERVICE_NAME              = "simple_zipkin_server"
	ZIPKIN_HTTP_ENDPOINT      = "http://127.0.0.1:9411/api/v1/spans"
	ZIPKIN_RECORDER_HOST_PORT = "127.0.0.1:9000"
)


func main(){
	c,err := credentials.NewClientTLSFromFile(os.Getenv("gopath") +"/src/Demo/grpc/conf/server.pem", "felix")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	}
	collector, err := zipkin.NewHTTPCollector(ZIPKIN_HTTP_ENDPOINT)
	if err != nil {
		log.Fatalf("zipkin.NewHTTPCollector err: %v", err)
	}

	recorder := zipkin.NewRecorder(collector, true, ZIPKIN_RECORDER_HOST_PORT, SERVICE_NAME)

	tracer, err := zipkin.NewTracer(
		recorder, zipkin.ClientServerSameSpan(false),
	)

	conn,err := grpc.Dial(":"+PORT,grpc.WithTransportCredentials(c),        grpc.WithUnaryInterceptor(
		otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads()),
	))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v",err)
	}
	defer  conn.Close()

	client := pb.NewSearchServiceClient(conn)
	resp,err := client.Search(context.Background(),&pb.SearchRequest{
		Request:"gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v",err)
	}
	log.Printf("resp: %s",resp.GetResponse())
}
