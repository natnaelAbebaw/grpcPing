package main

import (
	pb "backend/ping"
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func TestPingServer_Ping(t *testing.T){
	lis := bufconn.Listen(1024*1024)
	t.Cleanup(func() {
		lis.Close()
	})

	srv := grpc.NewServer()
	t.Cleanup(func() {
		srv.Stop() 
	})

	 pb.RegisterPingTestServer(srv, &PingServer{})

	 go func ()  {
		if err := srv.Serve(lis); err != nil{
			 log.Fatalf("Error on Srv.serve: %v",err)
		}
	 }()
	 
   dialer := func(context.Context, string) (net.Conn,error) {
		return lis.Dial()
	 }
	 conn, err := grpc.DialContext(context.Background(), "", grpc.WithContextDialer(dialer), grpc.WithInsecure())

	 t.Cleanup(func() {
		conn.Close()
	 })

	 if err != nil{
		 t.Fatalf("Error on grpc.DialContext %v",err)
	 }

	 client := pb.NewPingTestClient(conn)

	pingText := "natty"

	res, err := client.Ping(context.Background(),&pb.PingRequest{Ping: &pb.Ping{Text: pingText }})

	if err != nil{
		t.Fatalf("Error on client.Ping: %v",err)
	}

	if (res.Ping.Text != pingText){
		t.Fatalf("Unexpected value: %v", res.Ping.Text)
	}

}


