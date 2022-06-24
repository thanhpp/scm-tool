package configx

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Pass     string `mapstructure:"pass"`
	Database string `mapstructure:"database"`
}

func (d *DatabaseConfig) OverideWithEnv() {
	host := os.Getenv("DB_HOST")
	if len(host) != 0 {
		d.Host = host
	}
}

func (c DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Pass, c.Database)
}
