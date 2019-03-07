package interceptor

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"rbac-grpc-interceptor/proto"
	"strings"
)

func RBACUnaryServerChain(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	t := test.NewRBACClient(conn)

	FullMethod := strings.Split(info.FullMethod, "/")

	Sub := []string{"Admin", "Work"}

	tr, err := t.Enforce(context.Background(), &test.EnforceRep{
		Role:    Sub,
		Method:  FullMethod[1],
		Pattern: FullMethod[2],
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	if tr.Code == 0 {
		return nil, grpc.Errorf(codes.Unauthenticated, "Permission denied")
	} else {
		log.Printf("Permission to enter")
		return handler(ctx, req)
	}

}
