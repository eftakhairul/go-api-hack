package apis

import (
	"net/http"
	"strconv"

	"github.com/eftakhairul/go-api-hack/app/daos"
	"github.com/eftakhairul/go-api-hack/app/libs"
	models "github.com/eftakhairul/go-api-hack/app/models"
	"github.com/eftakhairul/go-api-hack/app/services"
	"github.com/gin-gonic/gin"
)

// GetUserOne godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func GetUserOne(c *gin.Context) {
	appContext := c.MustGet("appContext").(*libs.AppContext)
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	userService := services.NewUserService(daos.NewUserDAO(appContext))
	if user, err := userService.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		appContext.Logger.Error(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUser godoc
// @Summary Retrieves all user
// @Produce json
// @Success 200 {object} []models.User
// @Router /users [get]
// @Security ApiKeyAuth
func GetUser(c *gin.Context) {
	appContext := c.MustGet("appContext").(*libs.AppContext)

	userService := services.NewUserService(daos.NewUserDAO(appContext))
	if users, err := userService.GetAll(); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		appContext.Logger.Error(err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

// PostUser godoc
// @Summary Create user based request
// @Produce json
// @Success 200
// @Router /users [POST]
// @Security ApiKeyAuth
func PostUser(c *gin.Context) {
	appContext := c.MustGet("appContext").(*libs.AppContext)

	requestBody, ok := c.Get("body")
	user, ok := requestBody.(*models.User)
	if !ok {
		appContext.Logger.Error("error while processing input")
		c.JSON(http.StatusBadRequest, libs.NewErrorWrapper(http.StatusBadRequest, "input error"))
		return
	}

	userService := services.NewUserService(daos.NewUserDAO(appContext))
	if err := userService.Create(user); err != nil {
		appContext.Logger.Error("DB insertion failed")
		c.JSON(http.StatusInternalServerError, libs.NewErrorWrapper(http.StatusInternalServerError, "Internal server error"))
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
