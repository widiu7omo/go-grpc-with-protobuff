package main

import (
	"fmt"
	"share-lib/model"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	user := model.User{
		Id:   1,
		Name: "Santo",
		Age:  2,
	}
	//JSON
	userJson, err := protojson.Marshal(&user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Original => %#v \n", user)
	fmt.Printf("String => %#v \n", user.String())
	fmt.Printf("JSON => %#v \n", string(userJson))
}
