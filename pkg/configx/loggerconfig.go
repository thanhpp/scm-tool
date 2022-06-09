package configx

type LogConfig struct {
	Color      bool   `mapstructure:"color"`
	LoggerName string `mapstructure:"logger_name"`
	Level      string `mapstructure:"level"`
}
