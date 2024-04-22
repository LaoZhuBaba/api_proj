package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/laozhubaba/api_proj/internal/logger"
	"github.com/laozhubaba/api_proj/pkg/server"
	"github.com/laozhubaba/api_proj/pkg/server/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Controller struct {
	server.Logic
	logger.Logger
	proto.UnimplementedApiServer
	addr string
	port int
}

func NewControler(ctx context.Context, addr string, port int, logic server.Logic, logger server.Logger) (Controller, error) {
	return Controller{Logic: logic, addr: addr, port: port, Logger: logger}, nil
}

// GetUsers(*emptypb.Empty, Api_GetUsersServer) error
// GetUser(context.Context, *UserId) (*GetUserResponse, error)

func (c Controller) GetUser(ctx context.Context, u *proto.UserId) (*proto.GetUserResponse, error) {
	c.Log("in grcp controller for GetUser")
	id := u.Message
	person, _ := c.Logic.GetUser(int(id))
	return &proto.GetUserResponse{Name: person.Name, Address: person.Address}, nil
}
func (c Controller) GetUsers(ctx context.Context, _ *emptypb.Empty) (persons *proto.GetUsersResponse, err error) {
	users, _ := c.Logic.GetUsers()
	persons = &proto.GetUsersResponse{
		Users: []*proto.GetUserResponse{},
	}
	for _, user := range users {
		persons.Users = append(persons.Users, &proto.GetUserResponse{Name: user.Name, Address: user.Address})
	}
	return persons, nil
}

func (c Controller) Run() error {
	gsrv := grpc.NewServer()
	proto.RegisterApiServer(gsrv, c)
	reflection.Register(gsrv) // Required for troubleshooting with debug client
	lis, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", c.addr, c.port))
	gsrv.Serve(lis)
	return nil
}
