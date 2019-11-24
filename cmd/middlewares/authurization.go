package middlewares

import "github.com/gin-gonic/gin"

//AuthorizationMiddleware authorizes request
func AuthorizationMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		/*
			authHeader := context.GetHeader("Authorization")
			if len(authHeader) == 0 {
				libs.NewError(c, http.StatusUnauthorized, errors.New("Authorization is required Header"))
				context.Abort()
			}
			if authHeader != config.Config.ApiKey {
				libs.NewError(context, http.StatusUnauthorized, fmt.Errorf("this user isn't authorized to this operation: api_key=%s", authHeader))
				context.Abort()
			}
			context.Next()
		*/
	}
}
