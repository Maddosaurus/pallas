package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cm "github.com/Maddosaurus/pallas/lib"
	gw "github.com/Maddosaurus/pallas/proto/pallas"
)

func run() error {
	grpc_endpoint := cm.Getenv("PALLAS_GRPC_ENDPOINT", "server:50051")
	serv_endpoint := cm.Getenv("PALLAS_REST_SERVE_ENDPOINT", ":8081")

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterOtpHandlerFromEndpoint(ctx, mux, grpc_endpoint, opts)
	if err != nil {
		return err
	}

	log.Printf("REST API server listening at %v", serv_endpoint)

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(serv_endpoint, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
