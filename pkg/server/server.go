package server

import (
	"log"
	"net"

	"github.com/shivaraj-shanthaiah/godoc-user/pkg/handler"
	pb "github.com/shivaraj-shanthaiah/godoc-user/pkg/pb"
	"google.golang.org/grpc"
)

func NewgrpcServer(handlr *handler.UserHandler) {
	log.Println("Connecting to gRPC server")
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("Error creating listner on port 8082")
	}

	grp := grpc.NewServer()
	pb.RegisterUserServiceServer(grp, handlr)

	log.Printf("Listening on gRPC server on port 8082")
	if err := grp.Serve(lis); err != nil {
		log.Fatal("Error connecting to gRPC server")
	}
}
