package main

import (
	"context"
	"time"

	"github.com/laozhubaba/api_proj/internal/fakedb"
	"github.com/laozhubaba/api_proj/internal/logger"
	"github.com/laozhubaba/api_proj/pkg/server"

	"github.com/laozhubaba/api_proj/pkg/server/grpc"
)

func main() {
	ctx := context.Background()
	l := server.LoggerAdapter(logger.LogOutput)
	ds := fakedb.NewFakeDB(ctx, 200*time.Millisecond) // Second param is a fake latency value so we can test time-outs
	logic := server.NewSimpleLogic(l, ds)
	c, _ := grpc.NewControler(ctx, "localhost", 8080, logic, l)
	c.Run()
}
