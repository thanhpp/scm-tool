package scmcfg

import "github.com/thanhpp/scm/pkg/configx"

type MainConfig struct {
	HTTPServer configx.HTTPServerConfig `mapstructure:"http_server"`
	Logger     configx.LogConfig        `mapstructure:"logger"`
}
