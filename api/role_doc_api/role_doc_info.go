package role_doc_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
)

type RoleDocInfoResponse struct {
	IsPwd       bool   `json:"isPwd"`       // 是否开启密码
	RoleDocPwd  string `json:"roleDocPwd"`  // 角色文档的密码
	RolePwd     string `json:"rolePwd"`     // 角色的密码
	IsSee       bool   `json:"isSee"`       // 是否开启了试看
	FreeContent string `json:"freeContent"` // 文档的试看内容
}

// RoleDocInfoView 角色文档信息
// @Tags 角色文档管理
// @Summary 角色文档信息
// @Description 角色文档信息
// @Param token header string true "token"
// @Param data query RoleDocRequest true "参数"
// @Router /api/role_docs/info [get]
// @Produce json
// @Success 200 {object} res.Response{data=RoleDocInfoResponse}
func (RoleDocApi) RoleDocInfoView(c *gin.Context) {
	var cr RoleDocRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithError(err, c)
		return
	}

	var roleDoc models.RoleDocModel
	if err := global.DB.Preload("RoleModel").Take(&roleDoc, "role_id = ? and doc_id = ?", cr.RoleID, cr.DocID).Error; err != nil {
		res.FailWithMsg("文档不存在", c)
		return
	}
	response := RoleDocInfoResponse{
		RolePwd: roleDoc.RoleModel.Pwd,
	}
	if roleDoc.Pwd != nil {
		response.IsPwd = true
		response.RoleDocPwd = *roleDoc.Pwd
	}
	if roleDoc.FreeContent != nil {
		response.IsSee = true
		response.FreeContent = *roleDoc.FreeContent
	}
	res.OKWithData(response, c)
}
