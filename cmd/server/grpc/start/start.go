package start

import (
	"context"

	"github.com/laozhubaba/api_proj/cmd/server/common"
	"github.com/laozhubaba/api_proj/internal/logger"
	"github.com/laozhubaba/api_proj/pkg/server"

	"github.com/laozhubaba/api_proj/pkg/server/grpc"
)

func StartGrpc(ctx context.Context, l logger.Logger, ds server.DataStore) {
	l.Info("starting GRPC server on port %d", common.GrcpPort)
	logic := server.NewSimpleLogic(l, ds)
	c, _ := grpc.NewControler(ctx, "localhost", common.GrcpPort, logic, l)
	c.Run()
}
