package log_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/res"
	"time"
)

type LogRemoveRequest struct {
	IDList    []uint `json:"idList"`    // 可以传id列表删除
	StartTime string `json:"startTime"` // 年月日格式的开始时间
	EndTime   string `json:"endTime"`   // 年月日格式的结束时间
	UserID    uint   `json:"userID"`    // 根据用户删除日志
	IP        string `json:"ip"`        // 根据用户ip删除
}

// LogRemoveView 删除日志
// @Tags 日志管理
// @Summary 删除日志
// @Description 删除日志
// @Param data body LogRemoveRequest true "参数"
// @Param token header string true "token"
// @Router /api/logs [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (LogApi) LogRemoveView(c *gin.Context) {
	var cr LogRemoveRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	var logs []log_stash.LogModel
	if len(cr.IDList) > 0 {
		global.DB.Find(&logs, cr.IDList).Delete(&logs)
	} else if cr.UserID != 0 {
		global.DB.Find(&logs, "userID = ?", cr.UserID).Delete(&logs)
	} else if cr.IP != "" {
		global.DB.Find(&logs, "ip = ?", cr.IP).Delete(&logs)
	} else if cr.StartTime != "" && cr.EndTime != "" {
		_, startTimeErr := time.Parse("2006-01-02", cr.StartTime)
		_, endTimeErr := time.Parse("2006-01-02", cr.EndTime)
		if startTimeErr != nil {
			res.FailWithMsg("开始时间格式错误", c)
			return
		}
		if endTimeErr != nil {
			res.FailWithMsg("结束时间格式错误", c)
			return
		}
		global.DB.Find(&logs, "createdAt > date(?) and createdAt < date(?)", cr.StartTime, cr.EndTime).Delete(&logs)
	}
	res.OKWithMsg(fmt.Sprintf("共删除 %d 条日志", len(logs)), c)
}
