package main

import (
	"github.com/laozhubaba/api_proj/apps/server"
	"github.com/laozhubaba/api_proj/interfaces"
)

type Application struct {
	backend  interfaces.Backend
	frontend interfaces.Frontend
}

func main() {
	server.Serve()
}
