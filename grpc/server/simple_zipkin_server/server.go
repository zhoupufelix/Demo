package main

import (
	pb "Demo/grpc/proto"
	"context"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"log"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
	"google.golang.org/grpc"
)

type SearchService struct {

}

func(s *SearchService) Search(ctx context.Context,r *pb.SearchRequest)(*pb.SearchResponse,error){
	return &pb.SearchResponse{Response:r.GetRequest() + " Server"},nil
}

const(
	PORT = "9007"

	SERVICE_NAME = "simple_zipkin_server"
	ZIPKIN_HTTP_ENDPOINT = "http://127.0.0.1:9411/api/v1/spans"
	ZIPKIN_RECORDER_HOST_PORT = "127.0.0.1:9000"
)


func main() {
	collector, err := zipkin.NewHTTPCollector(ZIPKIN_HTTP_ENDPOINT)
	if err != nil {
		log.Fatalf("zipkin.NewHTTPCollector err:%v", err)
	}

	recorder := zipkin.NewRecorder(collector,true,ZIPKIN_RECORDER_HOST_PORT,SERVICE_NAME)

	tracer,err := zipkin.NewTracer(
		recorder,zipkin.ClientServerSameSpan(false),
	)
	if err != nil {
		log.Fatalf("zipkin.NewTracer err: %v", err)
	}
	c,err := credentials.NewServerTLSFromFile(os.Getenv("gopath") +"/src/Demo/grpc/conf/server.pem",
		os.Getenv("gopath") +"/src/Demo/grpc/conf/server.key")

	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.Creds(c),
		grpc_middleware.WithUnaryServerChain(
			otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads()),
		),
	}

	server := grpc.NewServer(opts...)
	pb.RegisterSearchServiceServer(server,&SearchService{})

	lis,err := net.Listen("tcp",":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v",err)
	}

	server.Serve(lis)

}
