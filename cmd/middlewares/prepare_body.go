package middlewares

//CORSMiddleware ...
func PrepareBodyMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		c.Next()
	}
}