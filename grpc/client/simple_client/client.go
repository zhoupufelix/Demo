package main

import (
	"google.golang.org/grpc"
	"log"
	pb "Demo/grpc/proto"
	"context"
)

const PORT = "9001"

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
