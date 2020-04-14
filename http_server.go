package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	pb "gorpc/proto"
	"net/http"
)

func StartHTTPProxy() error {
	mux := runtime.NewServeMux()

	ctx := context.Background()
	endpoint := "localhost:50051"
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := pb.RegisterDemoHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		return err
	}

	fmt.Println("http proxy started ...")
	return http.ListenAndServe("localhost:8080", mux)
}
