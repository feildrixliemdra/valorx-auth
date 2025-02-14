package router

import (
	"valorx-auth/internal/config"
	"valorx-auth/internal/handler"
	"valorx-auth/internal/middleware"

	"github.com/gin-gonic/gin"
)

type router struct {
	rtr     *gin.Engine
	handler *handler.Handler
	cfg     *config.Config
}

func NewRouter(rtr *gin.Engine, cfg *config.Config, handler *handler.Handler) Router {
	return &router{
		rtr,
		handler,
		cfg,
	}
}

type Router interface {
	Init()
}

func (r *router) Init() {
	r.rtr.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	userRouter := r.rtr.Group("/v1/users")

	// enable google auth depend on flag
	if r.cfg.Flag.IsGoogleAuthEnabled {
		authRouter := r.rtr.Group("/v1/auth")
		authRouter.GET("/google/login", r.handler.AuthHandler.GoogleLogin)
		authRouter.GET("/google/callback", r.handler.AuthHandler.GoogleCallback)

		// protect users route with JWT only if google auth is enabled
		userRouter.Use(middleware.JWTAuth(r.cfg.JWT.SecretKey))
	}

	userRouter.GET("/:id", r.handler.UserHandler.GetDetail)
	userRouter.POST("", r.handler.UserHandler.Create)
	userRouter.PUT("/:id", r.handler.UserHandler.Update)
	userRouter.DELETE("/:id", r.handler.UserHandler.Delete)

}
