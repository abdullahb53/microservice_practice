package main

import (
	"context"
	"log"
	"time"

	genpb "github.com/abdullahb53/microservices_practice/genpb/golang_service_adder"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	//
	dialCtx, cleanup := context.WithTimeout(ctx, time.Second*15)
	defer cleanup()

	conn, err := grpc.DialContext(dialCtx, "localhost:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := genpb.NewEventsClient(conn)

	newitem := &genpb.NewItem{
		Id:       "ooooooooooooooo",
		ItemName: "aqnhmkghgerwqeda232323",
	}
	// create post
	ev, err := client.AddItem(ctx, newitem)

	if err != nil {
		log.Fatalf("failed to create event %v:", err)
	}
	log.Printf("Created event: [%s],Response Status:%s", newitem.Id, ev.GetResponse())
	log.Fatalf("done")
}
