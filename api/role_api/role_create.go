package role_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
)

type RoleCreateRequest struct {
	ID    uint   `json:"id"`
	Title string `json:"title" binding:"required,min=2,max=16"`
	Pwd   string `json:"pwd"`
}

// RoleCreateView 创建角色
// @Tags 角色管理
// @Summary 创建角色
// @Description 创建角色
// @Param data body RoleCreateRequest true "参数"
// @Param token header string true "token"
// @Router /api/roles [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (RoleApi) RoleCreateView(c *gin.Context) {
	var cr RoleCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, c)
		return
	}
	var role models.RoleModel
	if err := global.DB.Take(&role, "title = ?", cr.Title).Error; err == nil {
		res.FailWithMsg("角色名称已存在", c)
		return
	}
	global.DB.Create(&models.RoleModel{
		Title: cr.Title,
		Pwd:   cr.Pwd,
	})
	res.OKWithMsg("角色创建成功", c)
}
