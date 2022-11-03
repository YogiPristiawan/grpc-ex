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

var storage *models.ProductLists

func init() {
	storage = new(models.ProductLists)
	storage.Lists = make([]*models.Product, 0)
}

type ProductsServer struct {
	models.UnimplementedProductsServer
}

func main() {
	s := grpc.NewServer()
	var productSrv ProductsServer

	models.RegisterProductsServer(s, productSrv)

	l, err := net.Listen("tcp", fmt.Sprintf(":%s", config.PRODUCT_SERVICE_PORT))
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.PRODUCT_SERVICE_PORT, err)
	}

	log.Printf("RPC server listen on port %s", config.PRODUCT_SERVICE_PORT)

	log.Fatal(s.Serve(l))
}

func (ProductsServer) FindAll(ctx context.Context, void *empty.Empty) (*models.ProductLists, error) {
	return storage, nil
}

func (ProductsServer) GetByUserId(ctx context.Context, product *models.Product) (*models.ProductLists, error) {
	var pl models.ProductLists
	for _, val := range storage.Lists {
		if val.UserId == product.UserId {
			pl.Lists = append(pl.Lists, val)
		}
	}

	return &pl, nil
}

func (ProductsServer) Create(ctx context.Context, product *models.Product) (*models.Product, error) {
	storage.Lists = append(storage.Lists, product)
	return product, nil
}
