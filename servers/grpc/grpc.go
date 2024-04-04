package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/laozhubaba/api_proj/interfaces"
	"github.com/laozhubaba/api_proj/servers/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

// grpcServer implements the Frontend interface and the ApiServer interface (from proto)
type grpcServer struct {
	proto.UnimplementedApiServer
	backend interfaces.Backend
	addr    string
	port    int
}

func NewGrpcServer(ctx context.Context, addr string, port int, backend interfaces.Backend) (grpcServer, error) {
	return grpcServer{backend: backend, addr: addr, port: port}, nil
}

func (a grpcServer) GetUsers(ctx context.Context, _ *emptypb.Empty) (*proto.GetUsersResponse, error) {
	users := a.backend.GetAllUsers()
	return &proto.GetUsersResponse{Message: fmt.Sprintf("%v", users)}, nil
}

func (a grpcServer) GetUser(ctx context.Context, u *proto.UserId) (*proto.GetUserResponse, error) {
	id := u.Message
	user, _ := a.backend.GetUserById(int(id))
	return &proto.GetUserResponse{Message: fmt.Sprintf("%v", user)}, nil
}

func (a grpcServer) Run() error {
	gsrv := grpc.NewServer()
	proto.RegisterApiServer(gsrv, a)
	reflection.Register(gsrv)
	lis, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", a.addr, a.port))
	gsrv.Serve(lis)
	return nil
}
