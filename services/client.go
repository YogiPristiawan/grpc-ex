package main

import (
	"context"
	"encoding/json"
	"fmt"
	"grpc-ex/common/config"
	"grpc-ex/common/models"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// user := models.User{
	// 	Id:   1,
	// 	Name: "Yogi Pristiawan",
	// }

	u := userService()

	// res, err := u.Create(context.Background(), &user)
	// if err != nil {
	// 	log.Fatalf("error during creating user service %v", err)
	// }

	// r, err := json.Marshal(res)
	// if err != nil {
	// 	log.Fatal("error during json marshal")
	// }

	// fmt.Println(string(r))

	// fmt.Println("===========================")

	ul, err := u.FindAll(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("error during call find user %v", err)
	}

	r, err := json.Marshal(ul)
	if err != nil {
		log.Fatal("error during call json marshal")
	}

	fmt.Println("FIND ALL", string(r))

}

func productService() models.ProductsClient {
	port := fmt.Sprintf(":%s", config.PRODUCT_SERVICE_PORT)

	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could ot start grpc client", err)
	}

	return models.NewProductsClient(conn)
}

func userService() models.UsersClient {
	port := fmt.Sprintf(":%s", config.USER_SERVICE_PORT)

	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("could not start grpc client", err)
	}

	return models.NewUsersClient(conn)
}
