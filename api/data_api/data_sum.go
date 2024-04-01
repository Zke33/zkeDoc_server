package data_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
)

type DataSumResponse struct {
	UserCount int `json:"userCount"`
	DocCount  int `json:"docCount"`
	DiggCount int `json:"diggCount"`
	LookCount int `json:"lookCount"`
}

// DataSumApiView 首页的求和数据
// @Tags 数据统计
// @Summary 首页的求和数据
// @Description 首页的求和数据
// @Router /api/data/sum [get]
// @Produce json
// @Success 200 {object} res.Response{data=DataSumResponse}
func (DataApi) DataSumApiView(c *gin.Context) {
	var response DataSumResponse
	global.DB.Model(models.UserModel{}).Select("count(id)").Scan(&response.UserCount)
	global.DB.Model(models.DocModel{}).Select("count(id)").Scan(&response.DocCount)
	global.DB.Model(models.DocModel{}).Select("sum(lookCount)").Scan(&response.LookCount)
	global.DB.Model(models.DocModel{}).Select("sum(diggCount)").Scan(&response.DiggCount)
	res.OKWithData(response, c)
}
