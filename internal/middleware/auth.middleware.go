package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kodacampmain/koda-b9-gin/internal/dto"
	"github.com/kodacampmain/koda-b9-gin/pkg"
)

func CheckToken(c *gin.Context) {
	bearer := c.GetHeader("Authorization")
	if bearer == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Msg:     "please login first",
		})
		return
	}
	result := strings.Split(bearer, " ")
	if len(result) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Msg:     "invalid bearer token",
		})
		return
	}
	if result[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Msg:     "invalid bearer token",
		})
		return
	}
	var token pkg.JWTClaims
	err := token.DecodeToken(result[1])
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenInvalidIssuer) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
				Success: false,
				Msg:     "invalid token",
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Msg:     "terjadi kesalahan sistem",
		})
		return
	}
	c.Set("token", token)
	c.Next()
}

// rbac => role based access control
func AdminOnly(c *gin.Context) {
	token, exists := c.Get("token")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Msg:     "missing token",
		})
		return
	}
	// cek role nya
	t, ok := token.(pkg.JWTClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Msg:     "invalid claims",
		})
		return
	}

	if t.Role != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
			Success: false,
			Msg:     "no privilege",
		})
		return
	}
	c.Next()
}
