package nftcfg

import (
	"errors"
	"os"

	"github.com/thanhpp/scm/pkg/configx"
)

const (
	infuraPrjIDEnv     = "INFURA_PROJECT_ID"
	infuraPrjSecretEnv = "INFURA_PROJECT_SECRET"
)

type NFTServiceConfig struct {
	HTTPServer   configx.HTTPServerConfig `mapstructure:"http_server"`
	Logger       configx.LogConfig        `mapstructure:"logger"`
	Database     configx.DatabaseConfig   `mapstructure:"database"`
	InfuraConfig InfuraConfig             `mapstructure:"-"`
}

func NewNFTServiceConfig(configPath string) (*NFTServiceConfig, error) {
	cfg := new(NFTServiceConfig)

	if err := configx.ReadConfigFromFile(configPath, cfg); err != nil {
		return nil, err
	}

	infuraPrjID := os.Getenv(infuraPrjIDEnv)
	infuraPrjSecret := os.Getenv(infuraPrjSecretEnv)

	if infuraPrjID == "" || infuraPrjSecret == "" {
		return nil, errors.New("empty infura config")
	}

	cfg.InfuraConfig = InfuraConfig{
		ProjectID:     infuraPrjID,
		ProjectSecret: infuraPrjSecret,
	}

	return cfg, nil
}

type InfuraConfig struct {
	ProjectID     string
	ProjectSecret string
}
