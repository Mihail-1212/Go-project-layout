package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mihail-1212/go-project-layout/pkg/domain"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	auth := api.Group("/user")
	{
		auth.GET("/", h.getUserList)
	}
}

// getUserList godoc
// @Summary Return user list
// @Produce json
// @Success 200 {object} Users
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/user/ [get]
func (h *Handler) getUserList(c *gin.Context) {
	var users Users
	var err error

	users.Users, err = h.services.UserService.GetAllUsers()

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

type Users struct {
	Users []domain.User `json:"users"`
}
