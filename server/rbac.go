package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	gw "rbac-grpc-interceptor/proto"
	authentication "rbac-grpc-interceptor/rbac"
)

type server struct{}

var RBAC = authentication.New()

func (s *server) SetPermit(ctx context.Context, in *gw.PermitRep) (*gw.PermitRes, error) {

	RBACList := []*authentication.Permit{}
	for i := 0; i < len(in.Permits); i++ {
		Permit := in.Permits[i]
		if Permit != nil {
			RBACList = append(RBACList, &authentication.Permit{
				Permit.Role,
				Permit.Method,
				Permit.Pattern,
				Permit.Module,
			})
		}
	}

	RBAC.LoadModel(RBACList)

	return &gw.PermitRes{Code: 1, Msg: "Modules list:Successful call"}, nil
}

func (s *server) Modules(ctx context.Context, in *gw.RoleRep) (*gw.ModuleRes, error) {
	RBACInfo := []*gw.ModuleInfo{}
	Models := RBAC.GetModel()
	for i := 0; i < len(in.Role); i++ {
		modules := Models[in.Role[i]]
		if modules != nil {
			for t := 0; t < len(modules); t++ {
				module := gw.ModuleInfo{
					Module: modules[t],
				}

				RBACInfo = append(RBACInfo, &module)
			}
		}
	}
	return &gw.ModuleRes{Code: 1, Msg: "Modules list:Successful call", DataInfo: RBACInfo}, nil
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

	return &gw.EnforceRes{Code: int64(Code), Msg: "GRPC Enforce:Successful call"}, nil
}

var (
	h bool
	p string
	i string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")

	flag.StringVar(&p, "p", "50052", "this is rpc port ,default:50052")

	flag.StringVar(&i, "i", "0.0.0.0", "this is rpc ip ,default:0.0.0.0")

	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `rbac version: rbac/0.0.3
Usage: rbac [-h] [-v] [-p] [-i]

Options:
`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	lis, err := net.Listen("tcp", i+":"+p)

	if err != nil {
		log.Fatalf("Server role_manage %s : %v", p, err)
	}
	log.Printf("listen: %v ....", p)

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	gw.RegisterRBACServer(s, &server{})

	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
