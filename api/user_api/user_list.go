package user_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/models"
	"gvd_server/service/common/list"
	"gvd_server/service/common/res"
)

// UserListView 用户列表
// @Tags 用户管理
// @Summary 用户列表
// @Description 用户列表
// @Param data query models.Pagination   false  "查询参数"
// @Router /api/users [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.UserModel]}
func (UserApi) UserListView(c *gin.Context) {
	var cr models.Pagination
	c.ShouldBindQuery(&cr)
	_list, count, _ := list.QueryList(models.UserModel{}, list.Option{
		Pagination: cr,
		Likes:      []string{"nickName", "userName"},
		Preload:    []string{"RoleModel"},
	})
	res.OKWithList(_list, count, c)
}
