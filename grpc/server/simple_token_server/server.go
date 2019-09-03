package main

import (
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	pb "Demo/grpc/proto"
	"net"
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"Demo/library"
	"time"
	"Demo/pkg/e"
)

type SearchService struct {

}

func(s *SearchService)Search(ctx context.Context,r *pb.SearchRequest)(*pb.SearchResponse,error){
	return &pb.SearchResponse{Response:r.GetRequest()+" Server"},nil
}

const(
	PORT = "9004"
)

func JwtInterceptor(ctx context.Context, req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler)(interface{}, error){
	//获取rpc 元数据
	md,ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil,status.Errorf(codes.Unauthenticated,"JWT Token 失败")
	}

	var token string
	if value,ok := md["token"];ok {
		token = value[0]
	}

	//检验token是否合法
	claims,err := library.ParseToken(token)
	log.Println(req)

	if err != nil {
		return nil,status.Errorf(codes.Unauthenticated,e.GetMsg(e.ERROR_AUTH_CHECK_TOKEN_FAIL))
	}else if time.Now().Unix() > claims.ExpiresAt {
		return nil,status.Errorf(codes.Unauthenticated,e.GetMsg(e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT))
	}

	resp, err := handler(ctx, req)
	return resp, err
}


func main(){
	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			JwtInterceptor,
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