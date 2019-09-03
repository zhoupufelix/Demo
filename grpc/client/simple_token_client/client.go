package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "Demo/grpc/proto"
	"google.golang.org/grpc/credentials"
	"os"
)

const (
	PORT = "9006"
)

type Auth struct {
	Token string
}

func (a *Auth) GetRequestMetadata(ctx context.Context,uri ...string)(map[string]string,error){
	return map[string]string{"token":a.Token},nil
}

func (a *Auth) RequireTransportSecurity()bool{
	return true
}

func main(){
	c,err := credentials.NewClientTLSFromFile(os.Getenv("gopath") +"/src/Demo/grpc/conf/server.pem", "felix")
	if err != nil {
		log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
	}
	auth := Auth{
		Token:"yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImZlbGl4IiwicGFzc3dvcmQiOiJ0d3VzYTEyMyIsImV4cCI6MTU2NzUwNDAyNiwiaXNzIjoiYXBpIn0.L3O-Xlalli-hhp0FgAvCLMvrALVtCciicQ8Kv-VDAk4",
	}
	conn,err := grpc.Dial(":"+PORT,grpc.WithTransportCredentials(c),grpc.WithPerRPCCredentials(&auth))
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