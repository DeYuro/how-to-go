package main

import (
	"context"
	proto "github.com/how-to-go/grpc/generated/proto/v1"
	"github.com/how-to-go/grpc/server"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

//git submodule update --init
//from grpc folder
//protoc-gen-go v1.27.1
//protoc -I /usr/local/include -I /Users/droot/go/src/github.com/how-to-go/grpc/api/dummy-proto  --go_out=generated --go_opt=paths=source_relative --go-grpc_out=generated --go-grpc_opt=paths=source_relative proto/v1/foo.proto proto/v1/structures.proto
//Or use some generator
func main() {
	if err := app(); err != nil {
		log.WithError(err).Fatal("application failed with error")
	}
}

func app() error {
	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan error)

	go func() {
		errCh <- startGrpcServer(cancel)
	}()

	select {
	case <-ctx.Done():
		log.Info("Service shutdown by ctx.Done")
		return nil
	case err := <-errCh:
		return err
	}
}

func startGrpcServer(cancel context.CancelFunc) error {
	grpcServer := grpc.NewServer()

	proto.RegisterFooServiceServer(grpcServer, server.NewServer())

	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		return err
	}
	err = grpcServer.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
