package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"userservice/internal/dao/query"
	"userservice/internal/service"
	"userservice/pb"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("userservice:%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := os.Getenv("DB_CONNECTION_STRING")
	if dsn == "" {
		log.Fatal("DB_CONNECTION_STRING environment variable is not set")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	query := query.Use(db)

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(grpcServer, service.NewUserServiceServer(query))
	grpcServer.Serve(lis)
}
