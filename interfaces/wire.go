//go:build wireinject
// +build wireinject

package interfaces

import (
	"context"

	"github.com/google/wire"

	"mcp-server-demo/infrastructure/common/cache"
	"mcp-server-demo/infrastructure/common/db"
	"mcp-server-demo/infrastructure/common/logit"
	"mcp-server-demo/infrastructure/common/request"
	"mcp-server-demo/infrastructure/common/response"
	"mcp-server-demo/interfaces/http"
	"mcp-server-demo/interfaces/http/controller"
)

func NewApp(ctx context.Context) (*app, error) {
	panic(wire.Build(wire.NewSet(
		loadAppConf,
		logit.NewServiceLoggerConf,
		logit.NewServiceLogger,

		db.NewDB,
		cache.NewRedis,
		response.NewHTTPResponseWriter,
		request.NewRequest,

		controller.NewHealth,
		http.NewHTTPHandler,
		newHTTPServer,

		wire.Struct(new(app), "*"),
	)))
}
