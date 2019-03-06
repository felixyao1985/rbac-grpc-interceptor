package main

import (
	"rbac-grpc-interceptor/authentication"
	gw "rbac-grpc-interceptor/proto"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// 业务实现方法的容器
type server struct{}

var ModActList []*gw.RBACInfo
var RBAC = authentication.New()

func (s *server) ModAct(ctx context.Context, in *gw.Rep) ( *gw.ModActRes, error) {

	RBACInfo := []*gw.RBACInfo{}

	return &gw.ModActRes{Code:1,Msg:"GRPC Menu List调用成功",DataInfo:RBACInfo}, nil
}

func (s *server) Enforce(ctx context.Context, in *gw.EnforceRep) (*gw.EnforceRes, error) {
	fmt.Println(in.Sub)
	fmt.Println(in.Obj)
	fmt.Println(in.Act)
	Code := 0
	for i := 0; i < len(in.Sub); i++ {
		p := RBAC.Enforce(in.Sub[i], in.Obj, in.Act)
		if(p) {
			Code = 1
			break
		}
	}

	return &gw.EnforceRes{Code:int64(Code),Msg:"GRPC Enforce调用成功"}, nil
}

func main() {


	lis, err := net.Listen("tcp", "0.0.0.0:50052") //监听所有网卡50051端口的TCP连接

	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	log.Printf("listen: %v ....", 50052)

	//获取权限数据
	//暂时先写死
	ModActList = append(ModActList,&gw.RBACInfo{
		Sub:"Admin",
		Obj:"test.Menu",
		Act:"View",
	})

	ModActList = append(ModActList,&gw.RBACInfo{
		Sub:"Member",
		Obj:"test.Menu",
		Act:"View",
	})

	ModActList = append(ModActList,&gw.RBACInfo{
		Sub:"Admin",
		Obj:"test.Menu",
		Act:"Save",
	})


	RBACList := []*authentication.ModAct{}
	for i := 0; i < len(ModActList); i++ {
		item := ModActList[i]
		if item != nil {
			RBACList = append(RBACList, &authentication.ModAct{
				item.Sub,
				item.Obj,
				item.Act,
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
