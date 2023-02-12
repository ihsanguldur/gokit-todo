package grpc

import (
	"context"
	todoSvc "go-kit-todo/services/todo"
	"go-kit-todo/services/todo/transport"
	"go-kit-todo/services/todo/transport/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func decodeCreateTodoRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateTodoRequest)
	return transport.CreateTodoRequest{
		Todo: todoSvc.Todo{
			Content: req.Todo.Content,
			State:   req.Todo.Status,
			UserID:  uint(req.Todo.UserID),
		},
	}, nil
}

func encodeCreateTodoResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(transport.CreateTodoResponse)
	err := getError(res.Err)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTodoResponse{Message: res.Message, Error: res.Err.Error()}, nil
}

func decodeListTodoRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return transport.ListTodoRequest{}, nil
}

func encodeListTodoResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(transport.ListTodoResponse)
	err := getError(res.Err)
	if err != nil {
		return nil, err
	}

	var todos []*pb.Todo
	for _, v := range res.Todos {
		todos = append(todos, &pb.Todo{
			Content: v.Content,
			Status:  v.State,
			UserID:  uint64(v.UserID),
		})
	}

	return &pb.ListResponse{Todos: todos, Error: res.Err.Error()}, nil
}

func decodeUpdateTodoRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateRequest)
	return transport.UpdateTodoRequest{
		ID: uint(req.Id),
		Todo: todoSvc.Todo{
			Content: req.Todo.Content,
			State:   req.Todo.Status,
			UserID:  uint(req.Todo.UserID),
		},
	}, nil
}

func encodeUpdateTodoResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(transport.UpdateTodoResponse)
	err := getError(res.Err)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateResponse{Message: res.Message, Error: res.Err.Error()}, nil
}

func getError(err error) error {
	if err == nil {
		return nil
	}

	return status.Error(codes.Unknown, err.Error())
}
