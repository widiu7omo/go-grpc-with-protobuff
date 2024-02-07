package main

import (
	"context"
	"log"
	"net"
	"share-lib/config"
	"share-lib/model"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

type UsersServer struct {
	model.UnimplementedUsersServer
}

func (UsersServer) Register(_ context.Context, param *model.User) (*emptypb.Empty, error) {
	user := param
	localStorage.List = append(localStorage.List, user)
	log.Println("New user registered", user.String())
	return new(emptypb.Empty), nil
}
func (UsersServer) List(context.Context, *emptypb.Empty) (*model.UserList, error) {
	return localStorage, nil
}
func (UsersServer) Add(_ context.Context, newUser *model.User) (*emptypb.Empty, error) {
	localStorage.List = append(localStorage.List, newUser)
	log.Println("New user added", newUser.String())
	return new(emptypb.Empty), nil
}
func main() {
	srv := grpc.NewServer()
	var usrServer UsersServer
	model.RegisterUsersServer(srv, usrServer)
	log.Println("Starting RPC server at", config.ServerPort)
	listener, err := net.Listen("tcp", config.ServerPort)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.ServerPort, err)
	}
	log.Fatal(srv.Serve(listener))
}
