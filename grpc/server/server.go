package server

import (
	"context"
	proto "github.com/how-to-go/grpc/generated/proto/v1"
)

type Server struct {
}

func (s *Server) Reverse(ctx context.Context, request *proto.ReverseRequest) (*proto.ReverseResponse, error) {
	panic("implement me")
}

func (s *Server) GetBar(ctx context.Context, request *proto.GetBarRequest) (*proto.GetBarResponse, error) {
	panic("implement me")
}

func NewServer() *Server {
	return &Server{}
}
