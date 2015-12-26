package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthTokenMiddleware checks token in Authorization header before request
func AuthTokenMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		RequestToken := c.Request.Header.Get("Authorization")
		if RequestToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "No Authorization Token specified!"})
			c.Abort()
			return
		}

		err := CheckToken(RequestToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}

}
