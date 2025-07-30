package v1

import "github.com/gin-gonic/gin"

func NewRouter(r *gin.Engine) {
	r.GET("/accounts", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
}
