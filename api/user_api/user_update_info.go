package user_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils/jwts"
)

type UserUpdateInfoRequest struct {
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}

// UserUpdateInfoView 修改用户信息
// @Tags 用户管理
// @Summary 修改用户信息
// @Description 修改用户信息
// @Param data body UserUpdateInfoRequest true  "UserUpdateInfoRequest"
// @Param token header string true  "token"
// @Router /api/users_info [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserUpdateInfoView(c *gin.Context) {
	var cr UserUpdateInfoRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims, _ := _claims.(*jwts.CustomClaims)
	user, err := claims.GetUser()
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}

	if cr.Avatar != "" {
		var image models.ImageModel
		err = global.DB.Take(&image, "userID = ? and path =  ?", claims.UserID, cr.Avatar[1:]).Error
		if err != nil {
			res.FailWithMsg("用户头像不存在", c)
			return
		}
	}
	if !(cr.NickName == "" && cr.Avatar == "") {
		global.DB.Model(user).Updates(models.UserModel{
			Avatar:   cr.Avatar,
			NickName: cr.NickName,
		})
	}
	res.OKWithMsg("用户信息修改成功", c)
}
