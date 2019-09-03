package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "Demo/grpc/proto"
)

const PORT = "9004"

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
	conn,err := grpc.Dial(":"+PORT,grpc.WithInsecure())
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