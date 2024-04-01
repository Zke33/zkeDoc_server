package role_doc_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils"
	"gvd_server/utils/jwts"
)

type RoleDocTree struct {
	ID       uint          `json:"key"`
	Title    string        `json:"title"`
	Children []RoleDocTree `json:"children"`
	IsPwd    bool          `json:"isPwd"`  // 是否需要密码
	Unlock   bool          `json:"unlock"` // 是否解锁
	IsColl   bool          `json:"isColl"` // 是否收藏
	IsSee    bool          `json:"isSee"`  // 是否试看
}

type RoleDocTreeResponse struct {
	List []RoleDocTree `json:"list"`
}

// RoleDocTreeView 角色文档树
// @Tags 角色文档管理
// @Summary 角色文档树
// @Description 角色文档树
// @Param token header string true "token"
// @Router /api/role_docs [get]
// @Produce json
// @Success 200 {object} res.Response{data=RoleDocTreeResponse}
func (RoleDocApi) RoleDocTreeView(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, err := jwts.ParseToken(token)
	var roleID uint = 2 // 默认给一个访客角色
	if err == nil {
		roleID = claims.RoleID
	}
	var response = RoleDocTreeResponse{
		List: make([]RoleDocTree, 0),
	}
	var docIDList []uint
	var roleDocList []models.RoleDocModel
	global.DB.
		Preload("RoleModel").
		Preload("DocModel").
		Find(&roleDocList, "role_id = ?", roleID).Select("doc_id").Scan(&docIDList)
	if len(roleDocList) == 0 {
		res.OKWithData(response, c)
		return
	}

	// 查文档列表
	var docList []*models.DocModel
	global.DB.Find(&docList, docIDList)
	if len(docList) == 0 {
		res.OKWithData(response, c)
		return
	}
	// 按照key里面点的个数排序，并且把最小的点数返回
	minCount := models.SortDocByPotCount(docList)
	// 构造一个新的docList
	var docListPointer = new([]*models.DocModel)
	// 循环排序好之后的 docList
	for _, model := range docList {
		// 判断，它们的key的点数是不是和最小的那个点数一样，一样就放入根列表
		if models.GetByPotCount(model) == minCount {
			// 根文档
			*docListPointer = append(*docListPointer, model)
			continue
		}
		// 子文档，需要找他们的父文档
		insertDoc(docListPointer, model)
	}
	// 文档树转换
	var docPwdMap = map[uint]bool{}
	var docSeeMap = map[uint]bool{}
	var docCollMap = map[uint]bool{}
	var unLuckMap = map[uint]bool{}

	for _, model := range roleDocList {
		// 判断有密码
		if model.Pwd != nil && (*model.Pwd != "" || model.RoleModel.Pwd != "") {
			docPwdMap[model.DocID] = true
		}
		// 判断试看
		if model.FreeContent != nil {
			docSeeMap[model.DocID] = true
		}
	}

	// 判断这个人
	if claims != nil && claims.UserID != 0 {
		// 判断是否收藏了
		// 判断是否解锁了
		var docCollIDList []uint
		global.DB.Model(models.UserCollDocModel{}).Where("user_id = ?", claims.UserID).
			Select("doc_id").Scan(&docCollIDList)
		var userPwdDocIDList []uint
		global.DB.Model(models.UserPwdDocModel{}).Where("user_id = ?", claims.UserID).
			Select("doc_id").Scan(&userPwdDocIDList)
		for _, id := range docCollIDList {
			docCollMap[id] = true
		}
		for _, id := range userPwdDocIDList {
			unLuckMap[id] = true
		}
	}

	list := RoleDocTreeTransition(*docListPointer, docPwdMap, docSeeMap, docCollMap, unLuckMap)
	response.List = list
	res.OKWithData(response, c)
}

// RoleDocTreeTransition 角色文档树 转换为特定类型
func RoleDocTreeTransition(docList []*models.DocModel, docPwdMap, docSeeMap, docCollMap, unLuckMap map[uint]bool) (list []RoleDocTree) {
	for _, model := range docList {
		children := RoleDocTreeTransition(model.Child, docPwdMap, docSeeMap, docCollMap, unLuckMap)
		if children == nil {
			children = make([]RoleDocTree, 0)
		}
		docTree := RoleDocTree{
			ID:       model.ID,
			Title:    model.Title,
			Children: children,
			IsPwd:    docPwdMap[model.ID],
			Unlock:   unLuckMap[model.ID],
			IsColl:   docCollMap[model.ID],
			IsSee:    docSeeMap[model.ID],
		}
		list = append(list, docTree)
	}
	return
}

// 找符合自己的父文档，并且插入进去
func insertDoc(docList *[]*models.DocModel, doc *models.DocModel) {
	// 把根文档的那个树一维化，通过最大字符前缀匹配，找到后面的key，最有可能匹配的key
	// 一维化
	oneDimensionalDocList := models.TreeByOneDimensional(*docList)
	// 通过最大前缀匹配找到这个model的key，对应应该放在哪一个对象上
	var keys []string
	for _, model := range oneDimensionalDocList {
		keys = append(keys, model.Key)
	}
	_, index := utils.FindMaxPrefix(doc.Key, keys)
	if index == -1 {
		// 没有满足的，那么就只能把它放到根文档上去了
		*docList = append(*docList, doc)
	} else {
		oneDimensionalDocList[index].Child = append(oneDimensionalDocList[index].Child, doc)
	}
}
