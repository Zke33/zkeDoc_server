package doc_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/res"
	"gvd_server/service/full_search_service"
)

// DocRemoveView 删除文档
// @Tags 文档管理
// @Summary 删除文档
// @Description 删除文档
// @Param id path int true "id"
// @Router /api/docs/{id} [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (DocApi) DocRemoveView(c *gin.Context) {
	var cr models.IDRequest
	c.ShouldBindUri(&cr)
	var doc models.DocModel
	if err := global.DB.Take(&doc, cr.ID).Error; err != nil {
		res.FailWithMsg("文档不存在", c)
		return
	}
	log := log_stash.NewAction(c)
	subDocList := models.FindAllSubDocList(doc)
	log.SetItemInfo("当前文档信息", map[string]any{
		"id":    doc.ID,
		"title": doc.Title,
	})
	// 先删所有子文档
	var docIdList []uint
	// 把自己的id放进去
	docIdList = append(docIdList, doc.ID)
	log.SetItemInfo("子文档数量", len(subDocList))
	for _, model := range subDocList {
		log.SetItemInfo("子文档信息", map[string]any{
			"id":    model.ID,
			"title": model.Title,
		})
		docIdList = append(docIdList, model.ID)
	}
	// 角色-文档表
	var roleDocList []models.RoleDocModel
	global.DB.Find(&roleDocList, "doc_id in ?", docIdList).Delete(&roleDocList)
	log.SetItemInfo("关联角色-文档数量", len(roleDocList))
	// 用户-收藏文档表
	var userDocList []models.UserCollDocModel
	global.DB.Find(&userDocList, "doc_id in ?", docIdList).Delete(&userDocList)
	log.SetItemInfo("关联用户-收藏文档数量", len(userDocList))
	// 用户-密码-文档表
	var userPwdDocList []models.UserPwdDocModel
	global.DB.Find(&userPwdDocList, "doc_id in ?", docIdList).Delete(&userPwdDocList)
	log.SetItemInfo("关联用户-密码-文档数量", len(userPwdDocList))
	// 删文档
	subDocList = append(subDocList, doc)
	if err := global.DB.Delete(&subDocList).Error; err != nil {
		log.SetItemErr("删除失败", err.Error())
		log.Error("文档删除失败")
		res.FailWithMsg("文档删除失败", c)
		return
	}
	go full_search_service.FullSearchDelete(doc.ID)
	log.Info(fmt.Sprintf("文档删除成功--%s", doc.Title))
	res.OKWithMsg(fmt.Sprintf("删除文档成功 共删除 %d 篇文档", len(subDocList)), c)
}
