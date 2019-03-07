package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"rbac-grpc-interceptor/examples/proto"
)

func main() {

	conn, err := grpc.Dial("0.0.0.0:50050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	t := proto.NewMenuClient(conn)

	tr, err := t.View(context.Background(), &proto.RepMenuView{ID: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Server Return : %s", tr.Msg)
}
