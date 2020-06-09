package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"m4-im/pkg/e"
	"m4-im/pkg/response"
	"m4-im/pkg/util"
	"net/http"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int = e.SUCCESS
		rsp := response.NewResponseBuilder(c)
		token := util.ParseBearerHeader(c.GetHeader("Authorization"))
		if token == "" {
			code = e.ErrInvalidBearerAuthParams
		} else {
			token, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ErrAuthCheckTokenTimeout
				default:
					code = e.ErrAuthCheckTokenFail
				}
			} else {
				c.Set("uid", token.Id)
			}
		}
		if code != e.SUCCESS {
			rsp.Code = code
			rsp.Response(http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
