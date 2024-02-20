package main

import (
	"crud_grpc/client/controler"
	"crud_grpc/client/routes"
	pb "crud_grpc/crud_grpc_proto"
	

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var Conn *grpc.ClientConn
var Instance pb.CrudClient

func main() {
	Conn, _ = grpc.Dial("localhost:8080", grpc.WithInsecure())
	Instance = pb.NewCrudClient(Conn)
	r := gin.Default()
	routes.Approutes(r)
	controler.GetInstance(Instance)
	
	r.Run(":6000")
}
