package main

import (
	"context"
	"crud_grpc/constants"
	"crud_grpc/controller"
	"fmt"
	"net"

	pb "crud_grpc/crud_grpc_proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var Client *mongo.Client


func main() {
	ctx := context.Background()
	connection := options.Client().ApplyURI(constants.ConnectionString)
	Client, _ = mongo.Connect(ctx, connection)
	Collections := Client.Database(constants.DbName).Collection("crud")
		controller.Setcollection(Collections)
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterCrudServer(s,&controller.Crud{})
	fmt.Println("server listening on port :8080")
	err1 := s.Serve(lis)
	
	if err1 != nil {
	
		fmt.Printf("failed to serve")
		return
	}

}
