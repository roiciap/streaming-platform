package main

import (
	"context"
	"log"
	"net"

	jwt_auth "github.com/roiciap/streaming-platform/be/go/proto/jwt-auth"
	"google.golang.org/grpc"
)

type myJwtAuth struct {
	jwt_auth.UnimplementedJWTServiceServer
}

func (s myJwtAuth) GenerateToken(ctx context.Context, req *jwt_auth.JWTTokenRequest) (*jwt_auth.JWTTokenResponse, error) {
	return &jwt_auth.JWTTokenResponse{
		Token: "test",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}

	serviceRegistrar := grpc.NewServer()
	service := &myJwtAuth{}

	jwt_auth.RegisterJWTServiceServer(serviceRegistrar, service)
	err = serviceRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
