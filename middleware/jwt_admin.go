package middleware

import (
	"github.com/gin-gonic/gin"
	"gvd_server/service/common/res"
	"gvd_server/service/redis_service"
	"gvd_server/utils/jwts"
)

func JwtAdmin(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		res.FailWithMsg("未携带token", c)
		c.Abort()
		return
	}
	claims, err := jwts.ParseToken(token)
	if err != nil {
		res.FailWithMsg("token错误", c)
		c.Abort()
		return
	}
	if claims.RoleID != 1 {
		res.FailWithMsg("权限错误", c)
		c.Abort()
		return
	}
	if redis_service.CheckLogout(token) {
		res.FailWithMsg("token已失效", c)
		c.Abort()
		return
	}
	c.Set("claims", claims)
}
