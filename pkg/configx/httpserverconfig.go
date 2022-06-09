package configx

type HTTPServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
