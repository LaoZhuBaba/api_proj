package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/laozhubaba/api_proj/pkg/server"
	"github.com/laozhubaba/api_proj/pkg/server/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Controller struct {
	Logic  server.LogicStream
	logger server.Logger
	proto.UnimplementedApiServer
	addr string
	port int
}

func NewGrpcController(ctx context.Context, addr string, port int, logic server.LogicStream, logger server.Logger) (Controller, error) {
	return Controller{Logic: logic, addr: addr, port: port, logger: logger}, nil
}

func (c Controller) AddUser(ctx context.Context, msg *proto.GetUserResponse) (*emptypb.Empty, error) {
	c.logger.Info("in grcp controller for AddUser")
	person := server.Person{Name: msg.Name, Address: msg.Address}
	err := c.Logic.AddUser(person)
	if err != nil {
		c.logger.Info("error returned from AddUser: %v", err)
	}
	return nil, err
}

func (c Controller) GetUser(ctx context.Context, u *proto.UserId) (*proto.GetUserResponse, error) {
	c.logger.Info("in grcp controller for GetUser")
	id := u.Message
	person, _ := c.Logic.GetUser(int(id))
	return &proto.GetUserResponse{Name: person.Name, Address: person.Address}, nil
}

func (c Controller) GetUsers(ctx context.Context, _ *emptypb.Empty) (persons *proto.GetUsersResponse, err error) {
	c.logger.Info("in non-streaming grcp controller for GetUsers")
	users, _ := c.Logic.GetUsers()
	persons = &proto.GetUsersResponse{
		Users: []*proto.GetUserResponse{},
	}
	for _, user := range users {
		persons.Users = append(persons.Users, &proto.GetUserResponse{Name: user.Name, Address: user.Address})
	}
	return persons, nil
}

func (c Controller) GetUsersStream(_ *emptypb.Empty, stream proto.Api_GetUsersStreamServer) error {
	c.logger.Info("in streaming grcp controller for GetUsers")
	users, _ := c.Logic.GetUsers()
	for _, user := range users {
		msg := &proto.GetUserResponse{Name: user.Name, Address: user.Address}
		if err := stream.Send(msg); err != nil {
			c.logger.Error("error while reading stream")
			return err
		}
	}
	return nil
}

func (c Controller) Run() {
	gsrv := grpc.NewServer()
	proto.RegisterApiServer(gsrv, c)
	reflection.Register(gsrv) // Required for troubleshooting with debug client
	lis, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", c.addr, c.port))
	gsrv.Serve(lis)
}
