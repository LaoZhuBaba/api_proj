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
	l := logger.New()
	ds := fakedb.NewFakeDB(ctx, 2000)
	logic := server.NewSimpleLogic(l, ds)
	c := rest.NewRestController(ctx, l, logic)
	http.ListenAndServe("localhost:8080", c.Rtr)
}
