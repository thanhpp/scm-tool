package nftcfg

import (
	"errors"
	"os"

	"github.com/thanhpp/scm/pkg/configx"
)

const (
	infuraPrjIDEnv     = "INFURA_PROJECT_ID"
	infuraPrjSecretEnv = "INFURA_PROJECT_SECRET"
	nodeAPIURLEnv      = "NODE_API_URL"
	privateKeyEnv      = "PRIVATE_KEY"
)

type NFTServiceConfig struct {
	HTTPServer   configx.HTTPServerConfig `mapstructure:"http_server"`
	Logger       configx.LogConfig        `mapstructure:"logger"`
	Database     configx.DatabaseConfig   `mapstructure:"database"`
	InfuraConfig InfuraConfig             `mapstructure:"-"`
	NodeAPIURL   string                   `mapstructure:"-"`
	PrivateKey   string                   `mapstructure:"-"`
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

	apiURL := os.Getenv(nodeAPIURLEnv)
	privateKey := os.Getenv(privateKeyEnv)
	if apiURL == "" || privateKey == "" {
		return nil, errors.New("empty node api url or private key")
	}

	cfg.InfuraConfig = InfuraConfig{
		ProjectID:     infuraPrjID,
		ProjectSecret: infuraPrjSecret,
	}
	cfg.NodeAPIURL = apiURL
	cfg.PrivateKey = privateKey

	// overwrite db host
	dbHost := os.Getenv("DB_HOST")
	if len(dbHost) != 0 {
		cfg.Database.Host = dbHost
	}

	return cfg, nil
}

type InfuraConfig struct {
	ProjectID     string
	ProjectSecret string
}
