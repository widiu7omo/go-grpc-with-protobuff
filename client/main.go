package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"share-lib/config"
	"share-lib/model"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func serviceUser() model.UsersClient {
	port := config.ServerPort
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return model.NewUsersClient(conn)
}
func main() {
	user := model.User{
		Id:   3,
		Name: "Santo Wijoyo",
		Age:  32,
	}

	srvUser := serviceUser()
	fmt.Printf("\n %s> user test\n", strings.Repeat("=", 10))

	srvUser.Register(context.Background(), &user)
	res1, err := srvUser.List(context.Background(), new(emptypb.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))

}
