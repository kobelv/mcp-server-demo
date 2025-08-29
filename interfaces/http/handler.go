package http

import (
	"net/http"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"

	"mcp-server-demo/infrastructure/common/logit"
	"mcp-server-demo/infrastructure/common/response"
	"mcp-server-demo/interfaces/http/controller"
	"mcp-server-demo/interfaces/middleware"
)

func NewHTTPHandler(
	logger logit.LoggerInterface,
	orderC *controller.Order,
	h *controller.Health,
	response response.HTTPResponseInterface,
) http.Handler {
	r := gin.New()
	r.Use(middleware.Logger(logger, response), middleware.Recovery(logger, true))
	ginpprof.Wrap(r)
	r.GET("/health/liveness", h.Liveness)
	r.GET("/health/readiness", h.Readiness)
	r.POST("/order/placeorder", orderC.PlaceOrder)
	return r
}
