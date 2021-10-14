package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mihail-1212/go-project-layout/internal/service"
	"github.com/mihail-1212/go-project-layout/pkg/logger"
)

var log, _ = logger.NewLoggerDev()

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {

	v1 := api.Group("/v1")
	{
		h.initUsersRoutes(v1)
	}
}

type UserContextIsEmptyError struct {
	Message string
}

func (e UserContextIsEmptyError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
