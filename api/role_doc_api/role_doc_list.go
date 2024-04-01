package role_doc_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
)

type DocTree struct {
	ID       uint      `json:"key"`
	Title    string    `json:"title"`
	Children []DocTree `json:"children"`
	IsPwd    bool      `json:"isPwd"` // 是否需要密码
	IsSee    bool      `json:"isSee"` // 是否试看
	Show     bool      `json:"show"`  // 角色是否可以看到文档
}

type RoleDocListResponse struct {
	List []DocTree `json:"list"`
}

// RoleDocListView 文档树列表
// @Tags 角色文档管理
// @Summary 文档树列表
// @Description 文档树列表
// @Param token header string true "token"
// @Param id path int true "id"
// @Router /api/role_docs/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=RoleDocListResponse}
func (RoleDocApi) RoleDocListView(c *gin.Context) {
	var id models.IDRequest
	if err := c.ShouldBindUri(&id); err != nil || id.ID == 0 {
		res.FailWithMsg("参数错误", c)
		return
	}
	var roleDocList []models.RoleDocModel
	global.DB.
		Preload("RoleModel").
		Preload("DocModel").
		Find(&roleDocList, "role_id = ?", id.ID)
	// 把所有文档给查出来
	// 角色-密码
	// 角色-试看
	tree := models.DocTree(nil)
	// 判断哪些文档是有密码的
	var docPwdMap = map[uint]bool{}
	var docSeeMap = map[uint]bool{}
	var docIDMap = map[uint]bool{}
	for _, model := range roleDocList {
		// 判断有密码
		if model.Pwd != nil && (*model.Pwd != "" || model.RoleModel.Pwd != "") {
			docPwdMap[model.DocID] = true
		}
		// 判断试看
		if model.FreeContent != nil {
			docSeeMap[model.DocID] = true
		}
		docIDMap[model.DocID] = true
	}
	// 判断哪些文档是有试看的
	list := DocTreeTransition(tree, docPwdMap, docSeeMap, docIDMap)
	res.OKWithData(RoleDocListResponse{
		List: list,
	}, c)
}

// DocTreeTransition 文档树转换为特定类型
func DocTreeTransition(docList []*models.DocModel, docPwdMap, docSeeMap, docIDMap map[uint]bool) (list []DocTree) {
	for _, model := range docList {
		children := DocTreeTransition(model.Child, docPwdMap, docSeeMap, docIDMap)
		if children == nil {
			children = make([]DocTree, 0)
		}
		docTree := DocTree{
			ID:       model.ID,
			Title:    model.Title,
			Children: children,
			IsPwd:    docPwdMap[model.ID],
			IsSee:    docSeeMap[model.ID],
			Show:     docIDMap[model.ID],
		}
		list = append(list, docTree)
	}
	return
}
