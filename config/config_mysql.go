package config

type Mysql struct {
	Host     string `yaml:"host"`     // 服务器地址:端口
	Port     string `yaml:"port"`     //端口
	Config   string `yaml:"config"`   // 高级配置
	DB       string `yaml:"db"`       // 数据库名
	Username string `yaml:"username"` // 数据库用户名
	Password string `yaml:"password"` // 数据库密码
	LogLevel string `yaml:"logLevel"` // 是否开启Gorm全局日志
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DB + "?" + m.Config
}

func (m *Mysql) GetLogLevel() string {
	return m.LogLevel
}
