package autoload

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`                // redis地址
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`                // redis端口
	Password string `mapstructure:"password" json:"password" yaml:"password"`    // redis密码
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                      // redis数据库
	PoolSize int    `mapstructure:"pool-size" json:"pool-size" yaml:"pool-size"` // 连接池大小
}
