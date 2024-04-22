package main

import (
	"context"
	"net/http"

	"github.com/laozhubaba/api_proj/internal/fakedb"
	"github.com/laozhubaba/api_proj/internal/logger"
	"github.com/laozhubaba/api_proj/pkg/server"
	"github.com/laozhubaba/api_proj/pkg/server/rest"
)

func main() {
	ctx := context.Background()
	l := server.LoggerAdapter(logger.LogOutput)
	ds := fakedb.NewFakeDB(ctx)
	logic := server.NewSimpleLogic(l, ds)
	c := rest.NewRestController(ctx, l, logic)
	http.ListenAndServe("localhost:8080", c.Rtr)
}
