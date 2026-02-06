package middleware

import "github.com/gin-gonic/gin"

func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")

		for _, r := range roles {
			if role == r {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(403, gin.H{"error": "forbidden"})
	}
}
