package bootstrap

import (
	"valorx-auth/internal/config"
	"valorx-auth/internal/handler"
	"valorx-auth/internal/router"

	"github.com/gin-gonic/gin"
)

func InitiateGinRouter(cfg *config.Config, handler *handler.Handler) *gin.Engine {
	r := gin.Default()

	//init router
	route := router.NewRouter(r, cfg, handler)
	route.Init()

	gin.SetMode(cfg.App.ReleaseMode)

	return r
}
