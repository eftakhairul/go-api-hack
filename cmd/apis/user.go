package apis

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/eftakhairul/go-api-hack/cmd/daos"
	"github.com/eftakhairul/go-api-hack/cmd/libs"
	"github.com/eftakhairul/go-api-hack/cmd/services"
	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func GetUser(c *gin.Context) {
	appContext := c.MustGet("appContext").(*libs.AppContext)
	s := services.NewUserService(daos.NewUserDAO(appContext))
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// PostUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users [POST]
// @Security ApiKeyAuth
func PostUser(c *gin.Context) {
	appContext := c.MustGet("appContext").(*libs.AppContext)
	appContext.Logger.Info("done")

	user, okay := c.Get("body")
	if okay {
		fmt.Println("%v", user)
	}
	fmt.Println("%v", user)
	c.JSON(http.StatusOK, gin.H{
		"status": "okay",
	})
}
