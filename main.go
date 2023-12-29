package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"grpc/book"
	"grpc/config"
	"grpc/proto"
	"log"
	"net"
)

func main() {
	config.LoadVariables()

	db, err := sql.Open("mysql", config.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", config.APIPort))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	var opts []grpc.ServerOption

	server := grpc.NewServer(opts...)

	proto.RegisterBookServiceServer(server, book.NewService(db))

	log.Printf("Server listening in port %d", config.APIPort)

	server.Serve(l)
}
