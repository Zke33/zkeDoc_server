package role_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
)

// RoleUpdateView 更新角色
// @Tags 角色管理
// @Summary 更新角色
// @Description 更新角色
// @Param data body RoleCreateRequest true "参数"
// @Param token header string true "token"
// @Router /api/roles [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (RoleApi) RoleUpdateView(c *gin.Context) {
	var cr RoleCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, c)
		return
	}
	if cr.ID == 0 {
		res.FailWithMsg("请选择更新的文档", c)
		return
	}
	var role models.RoleModel
	if err := global.DB.Take(&role, cr.ID).Error; err != nil {
		res.FailWithMsg("文档不存在", c)
		return
	}
	// 要用map去更新
	if err := global.DB.Model(&role).Updates(map[string]any{
		"title": cr.Title,
		"pwd":   cr.Pwd,
	}).Error; err != nil {
		global.Log.Error(err)
		res.FailWithMsg("更新失败", c)
		return
	}
	res.OKWithMsg("更新成功", c)
}
