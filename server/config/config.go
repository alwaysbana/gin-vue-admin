package config

// server 结构体
type Server struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// MySQL 结构体
type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Db       string `mapstructure:"db"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// jwt 结构体
type JWT struct {
	Secret     string `mapstructure:"secret"`
	Expiration int    `mapstructure:"expiration"`
	Issuer     string `mapstructure:"issuer"`
}

// 配置结构体
type Config struct {
	Server Server `mapstructure:"server"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Jwt    JWT    `mapstructure:"jwt"`
}
