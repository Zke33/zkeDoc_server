package flags

import (
	"github.com/sirupsen/logrus"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
)

func DB() {
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.UserModel{},
			&models.RoleModel{},
			&models.DocModel{},
			&models.UserCollDocModel{},
			&models.RoleDocModel{},
			&models.ImageModel{},
			&models.UserPwdDocModel{},
			&models.LoginModel{},
			&models.DocDataModel{},
			&log_stash.LogModel{},
		)
	if err != nil {
		logrus.Fatalf("数据库迁移失败 err:%s\n", err.Error())
	}
	logrus.Infof("数据库迁移成功\n")
}
