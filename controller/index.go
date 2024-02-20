package controller

import (
	"context"
	pb "crud_grpc/crud_grpc_proto"
	model "crud_grpc/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Crud struct {
	pb.UnimplementedCrudServer
}

var Collect *mongo.Collection

func Setcollection(coll *mongo.Collection) {
	Collect = coll

}

func (u *Crud) CreateUser(ctx context.Context, req *pb.UserDetails) (*pb.Response, error) {

	var user *model.UserDetails

	user = &model.UserDetails{
		Name:    req.Name,
		PhoneNo: req.PhoneNo,
		Address: req.Address,
		EmailId: req.EmailId,
	}
	fmt.Println(Collect)
	_, err := Collect.InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err.Error())
	}

	res := &pb.Response{
		Message: "success",
	}
	return res, nil
}

func (u *Crud) Listuser(ctx context.Context, rq *pb.Empty) (*pb.ListUserResponse, error) {
	//ctx = context.Background()
	var user []*pb.UserDetails
	cursor, _ := Collect.Find(ctx, bson.D{{}})

	for cursor.Next(ctx) {
		var User *pb.UserDetails
		err := cursor.Decode(&User)
		if err != nil {
			fmt.Printf(err.Error())
			return nil, nil
		}
		user = append(user, User)
	}
	res := &pb.ListUserResponse{
		ListUser: user,
	}
	return res, nil
}

func (u *Crud) GetUserByName(ctx context.Context, req *pb.Name) (*pb.UserDetails, error) {
	var us  *model.UserDetails
	var user *pb.UserDetails
	us= &model.UserDetails{}
	//ctx = context.Background()
	fmt.Println(req.UserName)
	us.Name = req.UserName

	name := us.Name
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := Collect.FindOne(ctx, query).Decode(&user)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, nil
	}
	return user, nil
}
func (u *Crud) DeleteUser(ctx context.Context, req *pb.Name) (*pb.Response, error) {
	//var us *model.UserDetails
	us:=&model.UserDetails{}
	us.Name = req.UserName
	//ctx = context.Background()
	name := us.Name
	filter := bson.D{bson.E{Key: "name", Value: name}}
	result, _ := Collect.DeleteOne(ctx, filter)
	if result.DeletedCount != 1 {
		return nil, nil
	}
	res := &pb.Response{
		Message: "success",
	}
	return res, nil
}
