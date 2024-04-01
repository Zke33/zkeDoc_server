package user_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils/pwd"
)

type UserUpdateRequest struct {
	ID       uint   `json:"id" binding:"required" label:"用户id"`
	Password string `json:"password"` // 密码
	NickName string `json:"nickName"` // 昵称
	RoleID   uint   `json:"roleID"`   // 角色id
}

// UserUpdateView 用户更改
// @Tags 用户管理
// @Summary 用户更改
// @Description 用户更改
// @Param data body UserUpdateRequest true  "UserUpdateRequest"
// @Param token header string true  "token"
// @Router /api/user [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserUpdateView(c *gin.Context) {
	var cr UserUpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, c)
		return
	}
	var user models.UserModel
	if err := global.DB.Take(&user, cr.ID).Error; err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	if cr.Password != "" {
		cr.Password = pwd.HashPwd(cr.Password)
	}
	if cr.RoleID != 0 {
		var role models.RoleModel
		if err := global.DB.Take(&role, cr.RoleID).Error; err != nil {
			res.FailWithMsg("角色不存在", c)
			return
		}
	}

	if err := global.DB.Model(&user).Updates(models.UserModel{
		Password: cr.Password,
		NickName: cr.NickName,
		RoleID:   cr.RoleID,
	}).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMsg("用户更新失败", c)
		return
	}
	res.OKWithMsg("用户更新成功", c)
}
