package main

import (
	"rbac-grpc-interceptor/unit"
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	gw "rbac-grpc-interceptor/proto"
	"go-study/lib/negroni"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

var (
	echoEndpoint = flag.String("hello_endpoint1", "localhost:50051", "endpoint of YourService1")
)

func run() error {


	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	opts = append(opts, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(unit.ValidateMiddleware())))
	err := gw.RegisterMenuHandlerFromEndpoint(ctx, gwmux, *echoEndpoint, opts)

	if err != nil {
		return err
	}


	mux := http.NewServeMux()
	mux.Handle("/",gwmux)
	mux.HandleFunc("/login", unit.Login)

	n := negroni.Classic()
	n.UseHandler(mux)

	log.Printf("listen: %v ....", 3344)
	n.Run(":3344")

	return nil
}


func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}