package cron_service

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/redis_service"
)

// SyncDocData 同步文档数据
func SyncDocData() {
	diggAll := redis_service.NewDocDigg().GetAll()
	lookAll := redis_service.NewDocLook().GetAll()

	var docList []models.DocModel
	global.DB.Find(&docList)
	for _, model := range docList {
		sID := fmt.Sprintf("%d", model.ID)
		digg := diggAll[sID]
		look := lookAll[sID]
		if digg == 0 && look == 0 {
			logrus.Infof("%s 无变化", model.Title)
			continue
		}

		newDigg := digg + model.DiggCount
		newLook := look + model.LookCount
		global.DB.Model(&model).Updates(models.DocModel{
			DiggCount: newDigg,
			LookCount: newLook,
		})
		logrus.Infof("%s 更新成功 digg + %d， look + %d", model.Title, digg, look)
	}

	redis_service.NewDocDigg().Clear()
	redis_service.NewDocLook().Clear()

}
