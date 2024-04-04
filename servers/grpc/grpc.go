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

// grpcServer implements the Frontend and the ApiServer interfaces
type grpcServer struct {
	proto.UnimplementedApiServer
	interfaces.Backend
	addr string
	port int
}

func NewGrpcServer(ctx context.Context, addr string, port int, backend interfaces.Backend) (grpcServer, error) {
	return grpcServer{Backend: backend, addr: addr, port: port}, nil
}

func (g grpcServer) GetUsers(ctx context.Context, _ *emptypb.Empty) (*proto.GetUsersResponse, error) {
	users := g.Backend.GetAllUsers()
	return &proto.GetUsersResponse{Message: fmt.Sprintf("%v", users)}, nil
}

func (g grpcServer) GetUser(ctx context.Context, u *proto.UserId) (*proto.GetUserResponse, error) {
	id := u.Message
	user, _ := g.Backend.GetUserById(int(id))
	return &proto.GetUserResponse{Message: fmt.Sprintf("%v", user)}, nil
}

func (g grpcServer) Run() error {
	gsrv := grpc.NewServer()
	proto.RegisterApiServer(gsrv, g)
	reflection.Register(gsrv)
	lis, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", g.addr, g.port))
	gsrv.Serve(lis)
	return nil
}
