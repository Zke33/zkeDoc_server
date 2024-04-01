package log_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/res"
)

// LogReadView 日志读取
// @Tags 日志管理
// @Summary 日志列表
// @Description 日志列表
// @Param data query models.IDRequest true "参数"
// @Param token header string true "token"
// @Router /api/logs/read [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (LogApi) LogReadView(c *gin.Context) {
	var cr models.IDRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	var log log_stash.LogModel
	if err := global.DB.Take(&log, cr.ID).Error; err != nil {
		res.FailWithMsg("日志不存在", c)
		return
	}
	if log.ReadStatus {
		res.OKWithMsg("日志读取成功", c)
		return
	}
	global.DB.Model(&log).Update("readStatus", true)
	res.OKWithMsg("日志读取成功", c)
}
