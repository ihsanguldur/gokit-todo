package main

import (
	"fmt"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/oklog/oklog/pkg/group"
	"go-kit-todo/services/todo"
	todoSvc "go-kit-todo/services/todo/implementation"
	"go-kit-todo/services/todo/repository"
	"go-kit-todo/services/todo/transport"
	grpctransport "go-kit-todo/services/todo/transport/grpc"
	"go-kit-todo/services/todo/transport/pb"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	port = ":5001"
	dsn  = "super:1234567@tcp(localhost:3306)/gokit-auth?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger, "svc", "todo", "ts", log.DefaultTimestampUTC, "clr", log.DefaultCaller)
	}

	_ = level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *gorm.DB
	{
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			_ = level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		_ = level.Info(logger).Log("msg", "database connection established")

		if err = db.AutoMigrate(&todo.Todo{}); err != nil {
			_ = level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		_ = level.Info(logger).Log("msg", "migrations completed")
	}

	var svc todo.Service
	{
		repo := repository.New(db, logger)
		svc = todoSvc.New(repo, logger)
	}

	var endpoints transport.Endpoints
	{
		endpoints = transport.MakeEndpoints(svc)
	}

	var (
		opTracing       = kitoc.GRPCServerTrace()
		serverOptions   = []kitgrpc.ServerOption{opTracing}
		todoService     = grpctransport.NewGrpcServer(endpoints, serverOptions, logger)
		grpcListener, _ = net.Listen("tcp", port)
		grpcServer      = grpc.NewServer()
	)

	var g group.Group
	{
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", port)
			pb.RegisterTodoServiceServer(grpcServer, todoService)
			return grpcServer.Serve(grpcListener)
		},
			func(error) {
				_ = grpcListener.Close()
			})
	}
	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		},
			func(error) {
				close(cancelInterrupt)
			})
	}

	level.Error(logger).Log("exit", g.Run())
}
