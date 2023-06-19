package grpcapi

import (
	"log"
	"net"

	"github.com/ddomd/maily/internal/mdb"
	pb "github.com/ddomd/maily/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedMailyServiceServer
	Port string
	DB   *mdb.MDB
}

func NewServer(port string, db *mdb.MDB) *Server {
	return &Server{
		Port: port,
		DB:   db,
	}
}

func (cfg *Server) Serve() {
	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMailyServiceServer(grpcServer, cfg)
	reflection.Register(grpcServer)
	log.Fatal(grpcServer.Serve(lis))
}
