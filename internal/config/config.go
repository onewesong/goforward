package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/onewesong/goforward/internal/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	// ListenAddr specifies the address to listen on with API.
	ListenAddr string `mapstructure:"listen_addr"`
	// LogFile specifies a file where logs will be written to. This value will
	// only be used if LogWay is set appropriately. By default, this value is
	// "console".
	LogFile string `ini:"log_file" json:"log_file"`
	// LogWay specifies the way logging is managed. Valid values are "console"
	// or "file". If "console" is used, logs will be printed to stdout. If
	// "file" is used, logs will be printed to LogFile. By default, this value
	// is "console".
	LogWay string `ini:"log_way" json:"log_way"`
	// LogLevel specifies the minimum log level. Valid values are "trace",
	// "debug", "info", "warn", and "error". By default, this value is "info".
	LogLevel string `ini:"log_level" json:"log_level"`
	// LogMaxDays specifies the maximum number of days to store log information
	// before deletion. This is only used if LogWay == "file". By default, this
	// value is 0.
	LogMaxDays int64 `ini:"log_max_days" json:"log_max_days"`
	// DisableLogColor disables log colors when LogWay == "console" when set to
	// true. By default, this value is false.
	DisableLogColor bool `ini:"disable_log_color" json:"disable_log_color"`
	// ForwardLinkList specifies a list of forward links.
	ForwardLinkList []string `mapstructure:"forward_links"`

	ForwardLinks models.ForwardLinks
}

func LoadCfg(cfgPath string) (*Config, error) {
	v := viper.New()
	if len(cfgPath) > 0 {
		v.SetConfigFile(cfgPath)
	} else {
		v.AddConfigPath(".")
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		v.AddConfigPath(exPath)
	}
	v.SetDefault("listenAddr", "0.0.0.0:5668")
	v.SetDefault("logWay", "console")
	v.SetDefault("logFile", "console")
	v.SetDefault("logLevel", "info")
	v.SetDefault("log_max_days", 3)
	v.SetDefault("disable_log_color", false)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Info("Config file not found, use default config")
		} else {
			// Config file was found but another error was produced
			return nil, fmt.Errorf("read config error: %s", err)
		}
	}
	conf := &Config{}
	if err := v.Unmarshal(conf); err != nil {
		return nil, fmt.Errorf("unmarshal config error: %s", err)
	}
	if err := checkCfg(conf); err != nil {
		return nil, fmt.Errorf("check config error: %s", err)
	}
	return conf, nil
}

func checkCfg(c *Config) error {
	if c == nil {
		return fmt.Errorf("config nil")
	}
	if len(c.ForwardLinkList) > 0 {
		for _, i := range c.ForwardLinkList {
			fl, err := models.NewForwardLink(i)
			if err != nil {
				return err
			}
			c.ForwardLinks = append(c.ForwardLinks, fl)
		}
	}
	return nil
}
