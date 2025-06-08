package grpc

import (
	"github.com/yusologia/go-core/v2/grpc"
	"service/internal/app/grpc/server"
)

func Register(srv *logiagrpc.GRPCServer) {
	srv.Register(
		&server.TestingServer{},
	)
}
