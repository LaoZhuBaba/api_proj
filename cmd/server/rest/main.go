package main

import (
	"context"

	"github.com/laozhubaba/api_proj/cmd/server/common"
	"github.com/laozhubaba/api_proj/cmd/server/rest/start"
)

func main() {
	common.Start(context.Background(), start.StartRest)
}
