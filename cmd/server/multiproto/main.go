package main

import (
	"context"
	"sync"

	"github.com/laozhubaba/api_proj/cmd/server/common"
	grcp_start "github.com/laozhubaba/api_proj/cmd/server/grpc/start"
	rest_start "github.com/laozhubaba/api_proj/cmd/server/rest/start"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go common.Start(context.Background(), grcp_start.StartGrpc)
	go common.Start(context.Background(), rest_start.StartRest)
	wg.Wait()
}
