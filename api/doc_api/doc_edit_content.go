package doc_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
)

// DocEditContentView 获取完整的正文
// @Tags 文档管理
// @Summary 获取完整的正文
// @Description 获取完整的正文
// @Param id path int true "id"
// @Router /api/docs/edit/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (DocApi) DocEditContentView(c *gin.Context) {
	var cr models.IDRequest
	if err := c.ShouldBindUri(&cr); err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	var doc models.DocModel
	if err := global.DB.Take(&doc, cr.ID).Error; err != nil {
		// 这个角色没有这个文档的权限
		res.FailWithMsg("文档不存在", c)
		return
	}
	res.OKWithData(doc.Content, c)
}
