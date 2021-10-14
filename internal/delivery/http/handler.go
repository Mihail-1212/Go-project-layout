package http

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mihail-1212/go-project-layout/internal/delivery/http/v1"
	"github.com/mihail-1212/go-project-layout/internal/service"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitAPI() *gin.Engine {
	router := gin.New()
	// Cors headers
	router.Use(corsMiddleware())
	router.Use(optionMiddleware)

	handlerV1 := v1.NewHandler(h.services)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}

	return router
}
