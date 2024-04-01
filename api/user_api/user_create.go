package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils/ip"
	"gvd_server/utils/pwd"
	"time"
)

type UserCreateRequest struct {
	UserName string `json:"userName" binding:"required" label:"用户名"` // 用户名
	Password string `json:"password" binding:"required"`             // 密码
	NickName string `json:"nickName"`                                // 昵称
	RoleID   uint   `json:"roleID" binding:"required"`               // 角色id
}

// UserCreateView 新增用户
// @Tags 用户管理
// @Summary 新增用户
// @Description 新增用户
// @Router /api/user [post]
// @Param data body UserCreateRequest true "UserCreateRequest"
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidError(err, &cr, c)
		return
	}
	var user models.UserModel
	if err := global.DB.Take(&user, "userName = ?", cr.UserName).Error; err == nil {
		res.FailWithMsg("用户名已存在", c)
		return
	}
	if cr.NickName == "" {
		// 昵称如果不存在，那么就要
		var maxID uint
		global.DB.Model(models.UserModel{}).Select("max(id)").Scan(&maxID)
		cr.NickName = fmt.Sprintf("用户_%d", maxID+1)
	}
	var role models.RoleModel
	if err := global.DB.Take(&role, cr.RoleID).Error; err != nil {
		res.FailWithMsg("角色不存在", c)
		return
	}

	_ip := c.ClientIP()
	if err := global.DB.Create(&models.UserModel{
		UserName:  cr.UserName,
		Password:  pwd.HashPwd(cr.Password),
		NickName:  cr.NickName,
		IP:        _ip,
		Addr:      ip.GetAddr(_ip),
		RoleID:    cr.RoleID,
		LastLogin: time.Now(),
	}).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMsg("用户创建失败", c)
		return
	}
	res.OKWithMsg("用户创建成功", c)
}
