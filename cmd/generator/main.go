package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	desk "name/generator/internal/app/generator"
	"name/generator/pkg/api/generator"
	"net"
	"net/http"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:8002", "gRPC server endpoint")
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	generator.RegisterGeneratorServer(grpcServer, desk.NewImplementation())

	var group errgroup.Group

	group.Go(func() error {
		return grpcServer.Serve(lis)
	})

	mux := runtime.NewServeMux()

	group.Go(func() error {
		if err1 := generator.RegisterGeneratorHandlerServer(ctx, mux, desk.NewImplementation()); err1 != nil {
			return err1
		}
		return http.ListenAndServe("localhost:8080", mux)
	})
	if err = group.Wait(); err != nil {
		log.Fatal(err)
	}
}
