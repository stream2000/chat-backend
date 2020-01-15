package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"m4-im/pkg/e"
	"m4-im/pkg/response"
	"m4-im/rbac/util"
	"net/http"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			token, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			} else {
				c.Set("uid", token.Id)
			}
		}

		if code != e.SUCCESS {
			rsp := response.NewJsonResponse(c)
			rsp.Code = code
			rsp.Data = data
			rsp.Response(http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next()
	}
}
