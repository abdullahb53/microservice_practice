package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	genpb "github.com/abdullahb53/microservices_practice/genpb/golang_service_adder"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

/*
- MONGO_INITDB_ROOT_USERNAME=citizix
  - MONGO_INITDB_ROOT_PASSWORD=S3cret
*/
func main() {
	// context
	_ = context.Background()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := os.Getenv("LISTEN_ADDR")

	// Listener
	listenAddr := fmt.Sprintf("%s:%s", addr, port)
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen:%s:%v", listenAddr, port)
	}

	GetMongoDBCli()

	// gRPC server
	eventServer := new(myEventServer)
	grpcServer := grpc.NewServer()
	genpb.RegisterEventsServer(grpcServer, eventServer)
	log.Println("starting grpc server")
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	log.Fatal(grpcServer.Serve(lis))

}

type myEventServer struct {
	genpb.UnimplementedEventsServer
}

var (
	client     *mongo.Client
	collection *mongo.Collection
)

func (m myEventServer) AddItem(ctx context.Context, item *genpb.NewItem) (*genpb.ResponseValue, error) {
	var response string

	// err := client.Connect(ctx)
	// if err != nil {
	// 	log.Fatalf("MongoDB connection failed:%v", err)
	// }
	// defer client.Disconnect(ctx)

	// res, err := collection.InsertOne(ctx, bson.M{item.Id: item.ItemName})
	// if err != nil {
	// 	log.Fatalf("mongodb insertion failed:%v", err)
	// 	response = "failed"

	// }
	// id := res.InsertedID
	// fmt.Println("inserted ID", id, "item ID:", item.Id, "item Name:", item.ItemName)
	response = "success"
	return &genpb.ResponseValue{
		Response: response,
	}, nil
}
