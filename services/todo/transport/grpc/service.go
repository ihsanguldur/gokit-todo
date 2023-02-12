package grpc

import (
	"context"
	kittransport "github.com/go-kit/kit/transport"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"go-kit-todo/services/todo/transport"
	"go-kit-todo/services/todo/transport/pb"
)

type grpcServer struct {
	pb.UnimplementedTodoServiceServer
	createTodo kitgrpc.Handler
	listTodo   kitgrpc.Handler
	updateTodo kitgrpc.Handler
	logger     log.Logger
}

func NewGrpcServer(endpoints transport.Endpoints, options []kitgrpc.ServerOption, logger log.Logger) pb.TodoServiceServer {
	errorLogger := kitgrpc.ServerErrorHandler(kittransport.NewLogErrorHandler(logger))
	options = append(options, errorLogger)

	return &grpcServer{
		createTodo: kitgrpc.NewServer(endpoints.Create, decodeCreateTodoRequest, encodeCreateTodoResponse, options...),
		listTodo:   kitgrpc.NewServer(endpoints.List, decodeListTodoRequest, encodeListTodoResponse, options...),
		updateTodo: kitgrpc.NewServer(endpoints.Update, decodeUpdateTodoRequest, encodeUpdateTodoResponse, options...),
		logger:     logger,
	}
}

func (g *grpcServer) CreateTodo(ctx context.Context, request *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	_, res, err := g.createTodo.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return res.(*pb.CreateTodoResponse), nil
}

func (g *grpcServer) ListTodo(ctx context.Context, empty *pb.Empty) (*pb.ListResponse, error) {
	_, res, err := g.listTodo.ServeGRPC(ctx, empty)
	if err != nil {
		return nil, err
	}

	return res.(*pb.ListResponse), nil
}

func (g *grpcServer) UpdateTodo(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	_, res, err := g.listTodo.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return res.(*pb.UpdateResponse), nil
}
