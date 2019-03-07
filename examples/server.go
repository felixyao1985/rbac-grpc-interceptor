package main

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"rbac-grpc-interceptor/examples/interceptor"
	gw "rbac-grpc-interceptor/examples/proto"
	gbac_proto "rbac-grpc-interceptor/proto"
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

func SetRBAC() {
	conn, err := grpc.Dial("0.0.0.0:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	t := gbac_proto.NewRBACClient(conn)

	_, err = t.SetPermit(context.Background(), &gbac_proto.PermitRep{
		Permits: []*gbac_proto.Permit{
			&gbac_proto.Permit{
				Role:    "Admin",
				Method:  "test.Menu",
				Pattern: "View",
				Module:  "User",
			},
			&gbac_proto.Permit{
				Role:    "Admin",
				Method:  "test.Menu",
				Pattern: "Save",
				Module:  "User",
			},
			&gbac_proto.Permit{
				Role:    "Member",
				Method:  "test.Menu",
				Pattern: "View",
				Module:  "Work",
			},
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	modules, err := t.Modules(context.Background(), &gbac_proto.RoleRep{Role: []string{"Admin"}})
	fmt.Println("get rbac module", modules)
}
func main() {
	SetRBAC()

	lis, err := net.Listen("tcp", "0.0.0.0:50050")

	if err != nil {
		log.Fatalf("Server role_manage 50051: %v", err)
	}
	log.Printf("listen: %v ....", 50050)

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
