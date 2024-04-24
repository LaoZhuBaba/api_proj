package start

import (
	"context"
	"fmt"
	"net/http"

	"github.com/laozhubaba/api_proj/cmd/server/common"
	"github.com/laozhubaba/api_proj/internal/logger"
	"github.com/laozhubaba/api_proj/pkg/server"
	"github.com/laozhubaba/api_proj/pkg/server/rest"
)

func StartRest(ctx context.Context, l logger.Logger, ds server.DataStore) {
	l.Info("starting REST server on port %d", common.RestPort)
	logic := server.NewSimpleLogic(l, ds)
	c := rest.NewRestController(ctx, l, logic)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", common.RestPort), c.Rtr)
}
