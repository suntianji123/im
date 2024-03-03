package main

import (
	"github.com/im/common/api"
	cp "github.com/im/common/app"
	"github.com/im/common/grpc"
	"github.com/im/signal/connector/app"
	"github.com/im/signal/connector/handler"
	"github.com/im/signal/connector/pkg/service"
	_ "github.com/im/signal/connector/rpc"
	pb "google.golang.org/grpc"
)

func main() {
	app.App = cp.NewBuilder().Build()

	app.App.GetTcpAcceptor().RegisterHandlers(handler.Handlers)

	srv := pb.NewServer()
	api.RegisterBatchServiceServer(srv, service.NewBatchServiceServerImpl())
	api.RegisterHeartBeartServiceServer(srv, service.NewHeartBeatServiceServerImpl())
	s := grpc.NewGrpcServer(app.App.GetBuilder().GetGrpcServerConfig().Port, srv)
	err := s.Init()
	if err != nil {
		panic(err)
	}
	if err = app.App.Start(); err != nil {
		panic(err)
	}
	s.Close()
}
