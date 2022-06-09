package configx

import (
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func ReadConfigFromFile(fileIn string, cfg interface{}) error {
	if cfg == nil {
		return errors.New("read nil config")
	}

	v := viper.New()

	v.AddConfigPath(filepath.Dir(fileIn))
	v.SetConfigFile(filepath.Base(fileIn))
	v.SetConfigType(strings.Replace(filepath.Ext(fileIn), ".", "", 1))

	if err := v.ReadInConfig(); err != nil {
		return errors.WithMessage(err, "read in cofig")
	}

	if err := v.Unmarshal(cfg); err != nil {
		return errors.WithMessage(err, "unmarshal config")
	}

	return nil
}
