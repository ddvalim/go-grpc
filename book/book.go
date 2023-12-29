package book

import (
	"context"
	"database/sql"
	"grpc/proto"
)

type BookServiceServer struct {
	proto.UnimplementedBookServiceServer
	db *sql.DB
}

func NewService(db *sql.DB) *BookServiceServer {
	return &BookServiceServer{
		db: db,
	}
}

func (s *BookServiceServer) Save(ctx context.Context, book *proto.Book) (*proto.BookResponse, error) {
	return &proto.BookResponse{
		Book: book,
	}, nil
}

func (s *BookServiceServer) Find(ctx context.Context, filter *proto.BookFilterResponse) (*proto.BookListResponse, error) {
	books := []*proto.Book{
		{Id: 1, Name: "name", Description: "description", Category: "category"},
	}

	return &proto.BookListResponse{Books: books}, nil
}
