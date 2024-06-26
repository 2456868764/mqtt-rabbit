package routes

import (
	"net/http"
	"strings"

	"bifromq_engine/pkg/api"
	"bifromq_engine/pkg/utils"
	"github.com/gin-gonic/gin"
)

var NoAuthenticationURLS = []string{
	"/api/auth/login",
	"/api/auth/captcha",
	"/ui/*",
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, url := range NoAuthenticationURLS {
			if strings.ToLower(c.Request.URL.Path) == url {
				c.Next()
				return
			}
			if strings.HasSuffix(url, "*") {
				if strings.HasPrefix(strings.ToLower(c.Request.URL.Path), strings.ToLower(url[:len(url)-1])) {
					c.Next()
					return
				}
			}
		}
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			api.Error(c, 10002, "请求未携带token，无权限访问")
			c.Abort()
			return
		}
		j := utils.NewJWT()
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				api.Error(c, 10002, "授权已过期")
				c.Abort()
				return
			}
			api.Error(c, 10002, err.Error())
			c.Abort()
			return
		}
		c.Set("uid", claims.UID)
		c.Next()
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
