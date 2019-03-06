package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	authentication "rbac-grpc-interceptor/interceptor"
	gw "rbac-grpc-interceptor/proto"
)

type server struct{}

var ModActList []*gw.RBACInfo
var RBAC = authentication.New()

func (s *server) Modules(ctx context.Context, in *gw.RoleRep) (*gw.ModuleRes, error) {
	RBACInfo := []*gw.RBACInfo{}
	return &gw.ModuleRes{Code: 1, Msg: "Modules list ：Successful call", DataInfo: RBACInfo}, nil
}

func (s *server) Enforce(ctx context.Context, in *gw.EnforceRep) (*gw.EnforceRes, error) {

	Code := 0
	for i := 0; i < len(in.Role); i++ {
		p := RBAC.Enforce(in.Role[i], in.Method, in.Pattern)
		if p {
			Code = 1
			break
		}
	}

	return &gw.EnforceRes{Code: int64(Code), Msg: "GRPC Enforce：Successful call"}, nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50052")

	if err != nil {
		log.Fatalf("Server role_manage 50052: %v", err)
	}
	log.Printf("listen: %v ....", 50052)

	ModActList = append(ModActList, &gw.RBACInfo{
		Role:    "Admin",
		Method:  "test.Menu",
		Pattern: "View",
		Module:  "User",
	})

	ModActList = append(ModActList, &gw.RBACInfo{
		Role:    "Member",
		Method:  "test.Menu",
		Pattern: "View",
		Module:  "User",
	})

	ModActList = append(ModActList, &gw.RBACInfo{
		Role:    "Admin",
		Method:  "test.Menu",
		Pattern: "Save",
		Module:  "Product",
	})

	RBACList := []*authentication.Permit{}
	for i := 0; i < len(ModActList); i++ {
		item := ModActList[i]
		if item != nil {
			RBACList = append(RBACList, &authentication.Permit{
				item.Role,
				item.Method,
				item.Pattern,
				item.Module,
			})
		}

	}

	RBAC.LoadModel(RBACList)
	fmt.Println(RBAC.GetModel())

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	gw.RegisterRBACServer(s, &server{})

	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
