package grcpclient

import (
	"context"
	"fmt"
	"log"

	"github.com/laozhubaba/api_proj/cmd/server/common"
	"github.com/laozhubaba/api_proj/pkg/server"
	"github.com/laozhubaba/api_proj/pkg/server/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func setupConn() (conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial(fmt.Sprintf("localhost:%d", common.GrcpPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("could not connect to GRPC server: %v", err)
		return nil, err
	}
	return conn, err
}

func AddUser(person server.Person) error {
	conn, err := setupConn()
	if err != nil {
		return err
	}
	defer conn.Close()
	client := proto.NewApiClient(conn)
	_, err = client.AddUser(context.Background(), &proto.GetUserResponse{Name: person.Name, Address: person.Address})
	if err != nil {
		log.Printf("client.AddUser() failed with: %v", err)
		return err
	}
	return nil
}

func GetUser(id int) (person *server.Person, err error) {
	conn, err := setupConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	person = &server.Person{}
	client := proto.NewApiClient(conn)
	resp, err := client.GetUser(context.Background(), &proto.UserId{Message: int64(id)})
	if err != nil {
		log.Printf("client.GetUser() failed with: %v", err)
		return nil, err
	}
	person.Name = resp.Name
	person.Address = resp.Address
	return person, nil
}

func GetAllUsers() (persons []server.Person, err error) {
	conn, err := setupConn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := proto.NewApiClient(conn)

	resp, err := client.GetUsers(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Printf("client.GetAllUsers() failed with: %v", err)
		return nil, err
	}
	for _, person := range resp.Users {
		persons = append(persons, server.Person{Name: person.Name, Address: person.Address})
	}
	return persons, nil
}
