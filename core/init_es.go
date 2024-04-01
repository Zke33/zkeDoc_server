package core

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gvd_server/global"
)

func InitEs() *elastic.Client {

	client, err := elastic.NewClient(
		elastic.SetURL(global.Config.Es.Addr),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(global.Config.Es.User, global.Config.Es.Password),
	)
	if err != nil {
		logrus.Fatalf(fmt.Sprintf("[%s] es连接失败, err:%s", global.Config.Es.Addr, err.Error()))
	}
	return client
}
