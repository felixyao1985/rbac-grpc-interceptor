package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"rbac-grpc-interceptor/proto"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	t := test.NewMenuClient(conn)

	tr, err := t.View(context.Background(), &test.RepMenuView{ID: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Server Return : %s", tr.Msg)
}
