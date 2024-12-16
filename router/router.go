package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pealan/golang-rest-api/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(deviceHandler *handler.DeviceHandler) *ServerHTTP {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//Swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	DeviceRoutes(router.Group("/device"), deviceHandler)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start() {
	err := sh.engine.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatal("Gin engine couldn't start")
	}
}
