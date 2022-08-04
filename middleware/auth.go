package middleware

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"user": "4939d01598c0206c4fe0bf785aedf93caac81e74dc0e72be0451336e2384a4997893432f856cf12076ad704daa20f4678b1c4fc5560290c63399c47076d47d47",
	})
}
