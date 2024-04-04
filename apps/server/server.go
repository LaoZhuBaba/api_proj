package server

import (
	"context"

	"github.com/laozhubaba/api_proj/backends/fake"
	"github.com/laozhubaba/api_proj/interfaces"

	//"github.com/laozhubaba/api_proj/servers/rest"
	"github.com/laozhubaba/api_proj/servers/grpc"
)

type Application struct {
	backend  interfaces.Backend
	frontend interfaces.Frontend
}

func Serve() {
	ctx := context.Background()
	app := Application{}
	app.backend, _ = fake.NewFakeDB(ctx)
	// get a REST Frontend interface
	// app.frontend, _ = rest.NewRestServer(ctx, "localhost", 3000, app.backend)
	// get a GRPC Frontend interface
	app.frontend, _ = grpc.NewGrpcServer(ctx, "localhost", 3000, app.backend)

	app.backend.AddUser("david")
	app.backend.AddUser("john")
	app.backend.AddUser("mary")
	app.frontend.Run()
}
