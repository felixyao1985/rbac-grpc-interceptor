package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"rbac-grpc-interceptor/interceptor"
	gw "rbac-grpc-interceptor/proto"
)

type server struct{}

func (s *server) List(ctx context.Context, in *gw.RepMenuList) (*gw.MenuRes, error) {

	Meuns := []*gw.MenuModel{}
	return &gw.MenuRes{Code: 1, Msg: "GRPC Menu List：Successful call", DataInfo: Meuns}, nil
}

func (s *server) View(ctx context.Context, in *gw.RepMenuView) (*gw.MenuRes, error) {

	Meuns := []*gw.MenuModel{}
	return &gw.MenuRes{Code: 1, Msg: "GRPC Menu View ：Successful call", DataInfo: Meuns}, nil
}

func (s *server) Save(ctx context.Context, in *gw.MenuModel) (*gw.MenuRes, error) {

	Meuns := []*gw.MenuModel{}

	return &gw.MenuRes{Code: 1, Msg: "GRPC Menu Save：Successful call", DataInfo: Meuns}, nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Server role_manage 50051: %v", err)
	}
	log.Printf("listen: %v ....", 50051)

	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			interceptor.RBACUnaryServerChain,
		),
	}

	s := grpc.NewServer(opts...)
	gw.RegisterMenuServer(s, &server{})

	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
