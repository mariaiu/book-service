package main

import (
	"fmt"
	"github.com/mariaiu/book/internal/config"
	"github.com/mariaiu/book/internal/handler"
	"github.com/mariaiu/book/internal/repository/mysql"
	pb "github.com/mariaiu/book/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Init config
	cfg, err := config.SetUpConfig()
	if err != nil {
		logger.Fatalf("error initializing config: %s", err.Error())
	}

	// Init repository
	repo, err := mysql.NewRepository(cfg)
	if err != nil {
		logger.Fatalf("failed to initialize db: %s", err.Error())
	}

	defer repo.Close()

	// Init handler
	h := handler.NewHandler(repo)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)); if err != nil {
		logger.Fatalln(err)
	}

	// Init GRPC server
	GrpcServer := grpc.NewServer()
	pb.RegisterBookServer(GrpcServer, h)

	logger.Println("start GRPC server")

	// Start GRPC-server
	go func() {
		if err = GrpcServer.Serve(listener); err != nil {
			logger.Fatal(err)
		}

		defer GrpcServer.Stop()
	}()

	// Graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	logger.Println("closing")
}
