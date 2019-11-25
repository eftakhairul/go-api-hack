package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/eftakhairul/go-api-hack/cmd/libs"
	"github.com/gin-gonic/gin"
)

func PrepareBodyMiddleware(input interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		appContext := context.MustGet("appContext").(*libs.AppContext)

		rawRequestBody, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			appContext.Logger.Error("A group of walrus emerges from the ocean")
			context.JSON(http.StatusBadRequest, libs.NewErrorWrapper(http.StatusBadRequest, "Error while reading the request"))
			return
		}

		decoder := json.NewDecoder(bytes.NewReader(rawRequestBody))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&input); err != nil {
			appContext.Logger.Error("Error while reading the request body")
			context.JSON(http.StatusBadRequest, libs.NewErrorWrapper(http.StatusBadRequest, fmt.Sprintf("Invalid Request Error: %v", err)))
			return
		}

		context.Set("body", input)
		context.Next()
	}
}
