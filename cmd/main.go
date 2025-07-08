package main

import (
	_ "em/cmd/docs"
	"em/internal/config"
	"em/internal/db"
	"em/internal/handler"
	"em/internal/repository"
	"em/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// @title People API
// @version 1.0
// @description API with enrich
// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.LoadConfig()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.Infof("Starting server on port: %s", cfg.Port)

	conn, err := db.DBConnect(cfg.DBURL)
	if err != nil {
		logrus.Fatalf("Error with db connecting: %v", err)
	}
	defer conn.Close()

	repo := repository.NewPersonRepo(conn)
	enricher := service.NewEnricher()
	h := handler.NewPersonHandler(repo,enricher)

	r := gin.Default()

	r.GET("/people", h.GetPeople)
	r.POST("/people",h.CreatePerson)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	r.Run(":" + cfg.Port)
}
