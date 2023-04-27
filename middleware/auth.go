package middleware

import (
	"context"
	"customer/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authString string

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")

		if auth == "" {
			c.Next()
			return
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]

		validate, err := service.JwtValidate(context.Background(), auth)
		if err != nil || !validate.Valid {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid token"})
			return
		}

		customClaim, _ := validate.Claims.(*service.JwtCustomClaim)

		ctx := context.WithValue(c.Request.Context(), "auth", customClaim)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func CtxValue(ctx context.Context) *service.JwtCustomClaim {
	raw, _ := ctx.Value("auth").(*service.JwtCustomClaim)
	return raw
}
