package cron_service

import (
	"github.com/robfig/cron/v3"
	"time"
)

func CornInit() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")

	Cron := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))

	// 每天的2点去同步数据
	Cron.AddFunc("0 0 2 * * ?", SyncDocData)

	Cron.Start()
}
