package config

// server 结构体
type server struct {
	port int    `mapstructure:"port"`
	mode string `mapstructure:"mode"`
}

// MySQL 结构体
type mysql struct {
	host     string `mapstructure:"host"`
	port     int    `mapstructure:"port"`
	db       string `mapstructure:"db"`
	user     string `mapstructure:"user"`
	password string `mapstructure:"password"`
}

// jwt 结构体
type jwt struct {
	secret     string `mapstructure:"secret"`
	expiration int    `mapstructure:"expiration"`
	issuer     string `mapstructure:"issuer"`
}

type Config struct {
	Server server `mapstructure:"server"`
	Mysql  mysql  `mapstructure:"mysql"`
	Jwt    jwt    `mapstructure:"jwt"`
}
