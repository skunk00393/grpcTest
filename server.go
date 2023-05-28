package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"toDoApp/protos/toDoItem"
	"toDoApp/protos/toDoList"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

type serverList struct {
	toDoList.ListServiceServer
}

type serverItem struct {
	toDoItem.ItemServiceServer
}

func (s *serverList) CreateNew(ctx context.Context, req *toDoList.NewListRequest) (*toDoList.NewListResponse, error) {
	return &toDoList.NewListResponse{
		Response: fmt.Sprint(CreateNewList(req.Name, req.CreatedOn)),
	}, nil
}

func (s *serverList) UpdateList(ctx context.Context, req *toDoList.UpdateListRequest) (*toDoList.UpdateListResponse, error) {
	return &toDoList.UpdateListResponse{
		Response: fmt.Sprint(UpdateListName(req.Id, req.NewName)),
	}, nil
}

func (s *serverList) DeleteList(ctx context.Context, req *toDoList.DeleteListRequest) (*toDoList.DeleteListResponse, error) {
	return &toDoList.DeleteListResponse{
		Response: fmt.Sprint(DeleteList(req.Id)),
	}, nil
}

func (s *serverList) GetList(ctx context.Context, req *toDoList.GetListRequest) (*toDoList.GetListResponse, error) {
	resp, _ := GetList(req.Id)
	return &toDoList.GetListResponse{
		Items: resp,
	}, nil
}

func (s *serverItem) DeleteItem(ctx context.Context, req *toDoItem.DeleteItemRequest) (*toDoItem.DeleteItemResponse, error) {
	return &toDoItem.DeleteItemResponse{
		Response: fmt.Sprint(DeleteItem(req.Id)),
	}, nil
}

func (s *serverItem) UpdateItem(ctx context.Context, req *toDoItem.UpdateItemRequest) (*toDoItem.UpdateItemResponse, error) {
	return &toDoItem.UpdateItemResponse{
		Response: fmt.Sprint(UpdateItem(req.Id, req.NewName)),
	}, nil
}

func (s *serverItem) MarkItem(ctx context.Context, req *toDoItem.MarkItemRequest) (*toDoItem.MarkItemResponse, error) {
	return &toDoItem.MarkItemResponse{
		Response: fmt.Sprint(MarkItem(req.Id)),
	}, nil
}

func (s *serverItem) CreateNewItem(ctx context.Context, req *toDoItem.NewItemRequest) (*toDoItem.NewItemResponse, error) {
	return &toDoItem.NewItemResponse{
		Response: fmt.Sprint(CreateItem(req.ListId, req.Item)),
	}, nil
}

var db *sql.DB

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	connStr := "postgres://postgres:postgres@localhost/toDoApp?sslmode=disable"
	db, _ = sql.Open("postgres", connStr)

	if errDb := db.Ping(); errDb != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	toDoList.RegisterListServiceServer(s, &serverList{})
	toDoItem.RegisterItemServiceServer(s, &serverItem{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
