package simple_server

import (
	"golang.org/x/net/context"
	pb "Demo/grpc/proto"
	"google.golang.org/grpc"
	"net"
	"log"
)

type SearchService struct {

}

func(s *SearchService)Search(ctx context.Context,r *pb.SearchRequest)(*pb.SearchResponse,error){
	return &pb.SearchResponse{Response:r.GetRequest()+" Server"},nil
}

const PORT = "9001"

func main(){
	//1.创建gRPC Server 对象
	server := grpc.NewServer()
	//2.将被调用的服务端接口注册到gRPC Server的内部注册中心
	pb.RegisterSearchServiceServer(server,&SearchService{})

	lis,err := net.Listen("tcp",":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v",err)
	}
	server.Serve(lis)
}
