package http

import (
	"valorx-auth/internal/bootstrap"
	"valorx-auth/internal/handler"
	"valorx-auth/internal/repository"
	"valorx-auth/internal/server"
	"valorx-auth/internal/service"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// Start function handler starting http listener
func Start() {
	var (
		cfg         = bootstrap.NewConfig()
		err         error
		postgreConn *gorm.DB
		mongoDBConn *mongo.Client
		repo        *repository.Repository
		hndler      *handler.Handler
		svc         *service.Service
		router      *gin.Engine
	)

	// bootstrap dependency
	bootstrap.SetJSONFormatter()

	if cfg.Postgre.IsEnabled {
		postgreConn, err = bootstrap.InitiatePostgreSQL(cfg)
		if err != nil {
			log.Fatalf("error connect to PostgreSQL | %v", err)
		}

	}
	repo = repository.InitiateRepository(repository.Option{
		DB: postgreConn,
	})
	svc = service.InitiateService(cfg, repo)
	hndler = handler.InitiateHandler(cfg, svc)

	router = bootstrap.InitiateGinRouter(cfg, hndler)

	serve := server.NewHTTPServer(cfg, router.Handler())
	defer serve.Done()

	if err := serve.Run(); err != nil {
		log.Fatalf("error running http server %v", err.Error())
	}

	return
}
