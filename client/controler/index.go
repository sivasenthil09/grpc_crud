package controler

import (
	"context"
	pb "crud_grpc/crud_grpc_proto"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

var ins pb.CrudClient

func GetInstance(in pb.CrudClient) {
	
	ins = in

	
}

func CreateUser(ctx *gin.Context) {


	var req pb.UserDetails
	err1:=ctx.ShouldBindJSON(&req)
	if err1!=nil{
		fmt.Println(err1.Error())
		return
	}
	response, err := ins.CreateUser(context.Background(), &req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})

}
func Listuser(ctx *gin.Context) {
	var req pb.Empty
	ctx.ShouldBindJSON(&req)
	response, err := ins.Listuser(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})

}

func GetUserByName(ctx *gin.Context) {
	var req pb.Name
	ctx.ShouldBindJSON(&req)
	response, err := ins.GetUserByName(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})

}
func DeleteUser(ctx *gin.Context) {
	var req pb.Name
	ctx.ShouldBindJSON(&req)
	response, err := ins.DeleteUser(context.Background(), &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": response})

}
