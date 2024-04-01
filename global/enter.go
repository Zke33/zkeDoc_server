package global

import (
	"github.com/cc14514/go-geoip2"
	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvd_server/config"
)

var (
	Log      *logrus.Logger
	Config   *config.Config
	DB       *gorm.DB
	Redis    *redis.Client
	ESClient *elastic.Client
	AddrDB   *geoip2.DBReader
)

const DocSplitSign = "---===---"
