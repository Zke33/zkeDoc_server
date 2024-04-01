package flags

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"gvd_server/global"
	"os/exec"
	"time"
)

func Dump() {
	mysql := global.Config.Mysql

	timer := time.Now().Format("20060102")

	sqlPath := fmt.Sprintf("%s_%s.sql", mysql.DB, timer)

	cmder := fmt.Sprintf("mysqldump -u%s -p%s %s > %s", mysql.Username, mysql.Password, mysql.DB, sqlPath)
	cmd := exec.Command("sh", "-c", cmder)

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		logrus.Errorln(err.Error(), stderr.String())
		return
	}
	logrus.Infof("sql文件 %s 导出成功", sqlPath)
}
