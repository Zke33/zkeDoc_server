package user_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/service/common/res"
	"gvd_server/service/redis_service"
	"gvd_server/utils/jwts"
	"time"
)

// UserLogoutView 登出
// @Tags 用户管理
// @Summary 登出
// @Description 登出
// @Param token header string true  "token"
// @Router /api/logout [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserLogoutView(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, _ := jwts.ParseToken(token)
	exp := claims.ExpiresAt
	diff := exp.Time.Sub(time.Now())
	if err := redis_service.Logout(token, diff); err != nil {
		global.Log.Error(err)
		return
	}
	res.OKWithMsg("用户注销成功", c)
}
