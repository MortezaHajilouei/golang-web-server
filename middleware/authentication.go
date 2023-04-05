package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/MortezaHajilouei/golang-web-server/config"
	"github.com/MortezaHajilouei/golang-web-server/repository"
	"github.com/MortezaHajilouei/golang-web-server/utils"
	"github.com/gin-gonic/gin"
)

func Session() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "JWT" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token != "" {
			sub, err := utils.ValidateToken(access_token, config.Config_.AccessTokenPublicKey)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
				return
			}
			str := fmt.Sprintf("%v", sub)
			user, err := repository.NewUserRepository(config.DB).GetUser(str)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
				return
			}

			ctx.Set("currentUser", user)
		}
		ctx.Next()
	}
}

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, hasSession := ctx.Get("currentUser")
		if user != nil && hasSession {
			ctx.Next()
		} else {
			ctx.JSON(403, "Authentication failed")
			ctx.Abort()
		}
	}
}
