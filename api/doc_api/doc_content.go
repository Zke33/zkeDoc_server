package doc_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/service/redis_service"
	"gvd_server/utils/jwts"
	"strings"
)

type DocContentResponse struct {
	Content   string `json:"content"`
	IsSee     bool   `json:"isSee"`     // 是否试看
	IsPwd     bool   `json:"isPwd"`     // 是否需要密码
	IsColl    bool   `json:"isColl"`    // 用户是否收藏
	LookCount int    `json:"lookCount"` // 浏览量
	DiggCount int    `json:"diggCount"` // 点赞量
	CollCount int    `json:"collCount"` // 收藏量
}

// DocContentView 文档内容
// @Tags 文档管理
// @Summary 文档内容
// @Description 文档内容
// @Param id path int true "id"
// @Router /api/docs/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{data=DocContentResponse}
func (DocApi) DocContentView(c *gin.Context) {
	var cr models.IDRequest
	if err := c.ShouldBindUri(&cr); err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	// 因为这个接口，不登录也能访问，所以需要在视图里面解析token
	token := c.Request.Header.Get("token")
	claims, err := jwts.ParseToken(token)
	var roleID uint = 2 // 访客
	if err == nil {
		// 说明登录了
		roleID = claims.RoleID
	}
	// 判断角色是否有这个文档的访问权限
	var roleDoc models.RoleDocModel
	if err = global.DB.
		Preload("DocModel.UserCollDocList").
		Preload("RoleModel").
		Take(&roleDoc, "role_id = ? and doc_id = ?", roleID, cr.ID).Error; err != nil {
		// 这个角色没有这个文档的权限
		res.FailWithMsg("文档鉴权失败", c)
		return
	}
	redis_service.NewDocLook().SetById(cr.ID)
	doc := roleDoc.DocModel
	docDigg := redis_service.NewDocDigg().GetById(doc.ID)
	docLook := redis_service.NewDocLook().GetById(doc.ID)
	var response = DocContentResponse{
		DiggCount: docDigg + doc.DiggCount,
		LookCount: docLook + doc.LookCount,
		CollCount: len(doc.UserCollDocList),
	}
	// IsSee 这个角色是不是对这个文档有试看
	// 正文分隔符
	// 文档里面的试看内容
	// 角色-文档的试看内容
	// 试看部分 优先级：角色文档试看  > 文档试看字段 > 文档按照特殊字符分隔的试看
	// 判断正文里面是不是有特殊分隔符
	isDocFree := strings.Contains(doc.Content, global.DocSplitSign)
	var freeContent string                                                 // 试看正文
	var content = strings.ReplaceAll(doc.Content, global.DocSplitSign, "") // 实际正文
	if isDocFree {
		_list := strings.Split(doc.Content, global.DocSplitSign)
		freeContent = _list[0]
	}
	// 通过判断角色文档的 FreeContent是不是nil，如果不是，那么就开启了试看
	if roleDoc.FreeContent != nil {
		// 如果 FreeContent 为空，对应的优先级也都为空，也算它开启试看，试看内容 空
		// 在前端设置试看的时候，判断一下，有没有对应的试看内容，没有就要提示给用户
		response.IsSee = true

		// 按照优先级去设置试看
		if doc.FreeContent != "" {
			freeContent = doc.FreeContent
		}
		if *roleDoc.FreeContent != "" {
			freeContent = *roleDoc.FreeContent
		}
	}
	// IsPwd  判断这个角色有没有密码
	if roleDoc.Pwd != nil && (*roleDoc.Pwd != "" || roleDoc.RoleModel.Pwd != "") {
		response.IsPwd = true
	}
	// IsColl
	if roleID != 2 {
		// 查用户是否收藏了文档
		var userDoc models.UserCollDocModel
		err = global.DB.Take(&userDoc, "doc_id = ? and user_id = ?", cr.ID, claims.UserID).Error
		if err == nil {
			response.IsColl = true
		}
		// 用户是否对这个文档免密
		var usePwd models.UserPwdDocModel
		err = global.DB.Take(&usePwd, "doc_id = ? and user_id = ?", cr.ID, claims.UserID).Error
		if err == nil {
			response.IsPwd = false
		}
	}
	// Content
	// 有密码。有试看  试看内容
	// 无密码，有试看  试看内容
	if response.IsSee {
		response.Content = freeContent
	}
	// 有密码，无试看  空
	if response.IsPwd && !response.IsSee {
		response.Content = ""
	}
	// 无密码，无试看  正文
	if !response.IsPwd && !response.IsSee {
		response.Content = content
	}
	res.OKWithData(response, c)
}
