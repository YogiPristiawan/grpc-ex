package main

import (
	"context"
	"fmt"
	"grpc-ex/common/config"
	"grpc-ex/common/models"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var storage *models.UserLists

func init() {
	storage = new(models.UserLists)
	storage.Lists = make([]*models.User, 0)
}

type UserServer struct {
	models.UnimplementedUsersServer
}

func main() {
	s := grpc.NewServer()
	var userServer UserServer

	models.RegisterUsersServer(s, userServer)

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", config.USER_SERVICE_PORT))
	if err != nil {
		log.Fatalf("could not listen on %s: %v", config.PRODUCT_SERVICE_PORT, err)
	}

	log.Printf("RPC server listen on port %s", config.USER_SERVICE_PORT)

	log.Fatal(s.Serve(l))
}

func (UserServer) Create(ctx context.Context, user *models.User) (*models.User, error) {
	storage.Lists = append(storage.Lists, user)
	return user, nil
}

func (UserServer) FindAll(ctx context.Context, void *empty.Empty) (*models.UserLists, error) {
	return storage, nil
}
