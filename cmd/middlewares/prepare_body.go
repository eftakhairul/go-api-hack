package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eftakhairul/go-api-hack/cmd/libs"
	"github.com/eftakhairul/go-api-hack/cmd/models"
	"github.com/gin-gonic/gin"
)

// PrepareBodyMiddleware converts json to struct and validate the input
func PrepareBodyMiddleware(input interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		appContext := c.MustGet("appContext").(*libs.AppContext)

		rawRequestBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			appContext.Logger.Error("A group of walrus emerges from the ocean")
			c.JSON(http.StatusBadRequest, libs.NewErrorWrapper(http.StatusBadRequest, "Error while reading the request"))
			c.Abort()
		}

		decoder := json.NewDecoder(bytes.NewReader(rawRequestBody))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(input); err != nil {
			appContext.Logger.Error("Error while reading the request body")
			c.JSON(http.StatusBadRequest, libs.NewErrorWrapper(http.StatusBadRequest, fmt.Sprintf("Invalid Request Error: %v", err)))
			c.Abort()
		}

		// request body validation
		requestBody, ok := input.(models.RequestBody)
		if ok {
			err := requestBody.Validate()
			if err != nil {
				appContext.Logger.Error("Validation error")
				c.JSON(http.StatusBadRequest, libs.NewErrorWrapper(http.StatusBadRequest, "valiation failed"))
				c.Abort()
			}
		}
        // *************************************************
		// Same way on demand sanitization is possible here
		// *************************************************

		c.Set("body", input)
		c.Next()
	}
}
