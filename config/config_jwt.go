package config

type Jwt struct {
	Expires int    `yaml:"expires"` // 过期时间
	Issuer  string `yaml:"issuer"`
	Secret  string `yaml:"secret"`
}
