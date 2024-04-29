package main

import (
	"context"

	"github.com/laozhubaba/api_proj/cmd/server/common"
	"github.com/laozhubaba/api_proj/cmd/server/rest/start"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	common.Start(ctx, start.StartRest)
	cancel()
}
