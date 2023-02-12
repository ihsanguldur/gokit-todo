package middleware

import todoSvc "go-kit-todo/services/todo"

type Middleware func(svc todoSvc.Service) todoSvc.Service
