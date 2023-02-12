package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	todoSvc "go-kit-todo/services/todo"
)

type Endpoints struct {
	Create endpoint.Endpoint
	List   endpoint.Endpoint
	Update endpoint.Endpoint
}

func MakeEndpoints(svc todoSvc.Service) Endpoints {
	return Endpoints{
		Create: makeCreateTodoEndpoint(svc),
		List:   makeListTodoEndpoint(svc),
		Update: makeUpdateTodoEndpoint(svc),
	}
}

func makeCreateTodoEndpoint(svc todoSvc.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateTodoRequest)
		message, err := svc.CreateTodo(req.Todo)
		return CreateTodoResponse{Message: message, Err: err}, nil
	}
}

func makeListTodoEndpoint(svc todoSvc.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(ListTodoRequest)
		todos, err := svc.ListTodo()
		return ListTodoResponse{Todos: todos, Err: err}, nil
	}
}

func makeUpdateTodoEndpoint(svc todoSvc.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateTodoRequest)
		message, err := svc.UpdateTodo(req.ID, req.Todo)
		return UpdateTodoResponse{Message: message, Err: err}, nil
	}
}
