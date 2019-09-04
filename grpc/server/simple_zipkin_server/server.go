package main

import (
	pb "Demo/grpc/proto"
	"context"

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


func main(){
	collector,err := 
}
