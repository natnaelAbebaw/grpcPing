package main

import (
	pb "backend/ping"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type PingServer struct {
	pb.UnimplementedPingTestServer
}

func (s *PingServer) Ping(ctx context.Context, ping *pb.PingRequest) (*pb.PingResponse,error) {
	return  &pb.PingResponse{ Ping: ping.Ping} ,nil
}


func main(){
	port := ":50051"
	lis,err := net.Listen("tcp",port)
	if err != nil{
		log.Fatalf("Failed to listen: %v",err)
	}
	 defer lis.Close()
	 server := grpc.NewServer()
	 pb.RegisterPingTestServer(server,&PingServer{})

	 if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	
}