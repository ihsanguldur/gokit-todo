package transport

import todoSvc "go-kit-todo/services/todo"

type CreateTodoRequest struct {
	Todo todoSvc.Todo `json:"todo"`
}

type CreateTodoResponse struct {
	Message string `json:"message"`
	Err     error  `json:"error,omitempty"`
}

func (r *CreateTodoResponse) Failed() error { return r.Err }
func (r *CreateTodoResponse) Error() error  { return r.Err }

type ListTodoRequest struct{}

type ListTodoResponse struct {
	Todos []todoSvc.Todo `json:"todos"`
	Err   error          `json:"error,omitempty"`
}

func (r *ListTodoResponse) Failed() error { return r.Err }
func (r *ListTodoResponse) Error() error  { return r.Err }

type UpdateTodoRequest struct {
	ID   uint         `json:"id"`
	Todo todoSvc.Todo `json:"todo"`
}

type UpdateTodoResponse struct {
	Message string `json:"message"`
	Err     error  `json:"err,omitempty"`
}

func (r *UpdateTodoResponse) Failed() error { return r.Err }
func (r *UpdateTodoResponse) Error() error  { return r.Err }
