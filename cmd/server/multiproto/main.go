package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/laozhubaba/api_proj/cmd/server/common"
	grcp_start "github.com/laozhubaba/api_proj/cmd/server/grpc/start"
	rest_start "github.com/laozhubaba/api_proj/cmd/server/rest/start"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(2)
	go func() {
		defer wg.Done()
		common.Start(ctx, grcp_start.StartGrpc)
	}()
	go func() {
		defer wg.Done()
		common.Start(ctx, rest_start.StartRest)
	}()
	wg.Wait()
	fmt.Printf("exiting at the end of main() with error: %v\n", ctx.Err())
}
