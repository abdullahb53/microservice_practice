package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	genpb "github.com/abdullahb53/microservices_practice/genpb/golang_service_adder"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

/*
- MONGO_INITDB_ROOT_USERNAME=citizix
  - MONGO_INITDB_ROOT_PASSWORD=S3cret
*/

var (
	itemsCollection *mongo.Collection
	client          *mongo.Client
)

func init() {
	// create new client. (*mongo.Client)
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://citizix:S3cret@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// try to connect client. -> (*mongo.Client)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("client.Connect(ctx) failed:,%v", err)
	}

	// ping to Mongo DB server.
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		// Can't connect to Mongo server
		log.Fatalf("Cannot connect mongo server:%v", err)
	}
	print("Connected to Mongo DB Server\n")

	// get database and collection.
	quickstartDatabase := client.Database("itemsDB")
	itemsCollection = quickstartDatabase.Collection("itemsCollection")

}

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

	// GetMongoDBCli()

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

// Add item func impl.
func (m myEventServer) AddItem(ctx context.Context, item *genpb.NewItem) (*genpb.ResponseValue, error) {
	var response string
	var err error
	//---------------------------------------------
	// DB ADD TO COLLECTION -> (itemsCollection)---
	// insert one item to mongo db collection/table.
	_, err = itemsCollection.InsertOne(ctx, bson.D{
		{Key: item.Id, Value: item.ItemName},
	})
	if err != nil {
		log.Fatalf("itemsCollection add failed(mongo db):%v", err)
	}
	fmt.Println("------------------------")
	fmt.Println("-mongodb insert success-")
	fmt.Println("------------------------")
	fmt.Println("item.Id:[", item.Id, "]\n", "item.ItemName:[", item.ItemName, "]")

	// response value.
	response = "success"
	return &genpb.ResponseValue{
		Response: response,
	}, nil
}
