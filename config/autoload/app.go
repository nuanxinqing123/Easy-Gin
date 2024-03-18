package autoload

type App struct {
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
	Mode string `mapstructure:"mode" json:"mode" yaml:"mode"`
}
