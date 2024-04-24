package common

import (
	"context"
	"time"

	"github.com/laozhubaba/api_proj/internal/fakedb"
	"github.com/laozhubaba/api_proj/internal/logger"
	"github.com/laozhubaba/api_proj/pkg/server"
)

const (
	GrcpPort = 8081
	RestPort = 8080
)

// Start sets up common components (logger and datastore) then launches the specified server)
func Start(ctx context.Context, server func(context.Context, logger.Logger, server.DataStore)) {
	l := logger.New()
	ds := fakedb.NewFakeDB(ctx, 200*time.Millisecond) // Second param is a fake latency value so we can test time-outs
	server(ctx, l, ds)
}
