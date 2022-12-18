package middleware

import (
	"dkgosql-merchant-service-v4/internals/util"
	"errors"
	"fmt"

	auth "github.com/dhananjayksharma/dkgo-auth/auth"
	"github.com/gin-gonic/gin"
)

func ValidateRefreshToken(c *gin.Context) error {
	tokenString := c.GetHeader("RefreshToken")
	if tokenString == "" {
		err := errors.New("request does not contain an access token")
		if err != nil {
			return &util.InternalServer{ErrMessage: err.Error()}
		}
	}
	err := auth.ValidateRefreshToken(tokenString)
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

func CheckAuth(c *gin.Context) error {
	tokenString := c.GetHeader("Token")
	if tokenString == "" {
		err := errors.New("request does not contain an access token")
		if err != nil {
			return &util.InternalServer{ErrMessage: err.Error()}
		}
	}
	err := auth.ValidateToken(tokenString)
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

func Claim(c *gin.Context) error {
	tokenString := c.GetHeader("Token")

	claims, err := auth.GetClaim(tokenString)
	fmt.Printf("Username %v", claims.Username)
	fmt.Printf("Email %v", claims.Email)
	if err != nil {
		return &util.InternalServer{ErrMessage: err.Error()}
	}
	return nil
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Token")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
