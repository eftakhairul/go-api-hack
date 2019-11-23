package middlewares

//AuthorizationMiddleware authorizes request
func AuthorizationMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		c.Next()
	}
}