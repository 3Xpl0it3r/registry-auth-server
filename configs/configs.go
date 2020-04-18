package configs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var defaultConfig *Configs

type Configs struct {
	Server struct {
		Domain  string `json:"domain"`
		Address string `json:"address"`
		Port    string `json:"port"`
	} `json:"server"`
	Token struct {
		Issuer     string `json:"issuer"`
		Expiration int64  `json:"expiration"`
	} `json:"token"`
	SecureModule bool `json:"secureModule"`
	Tls          struct {
		Cert string `json:"cert"`
		Key  string `json:"key"`
	} `json:"tls"`
}

func NewConfigs(cfg string) *Configs {
	if defaultConfig == nil {
		defaultConfig = &Configs{}
		if cfg == "" {
			cfg = "config"
		}
		if err := defaultConfig.initConfig(cfg); err != nil {
			logrus.Panicf("init Config Failed: %s\n", err.Error())
		}
	}
	return defaultConfig
}

func (c *Configs) initConfig(cfg string) error {
	if c == nil {
		return fmt.Errorf("config is not initilization")
	}
	viper.SetConfigName(cfg)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(c); err != nil {
		return fmt.Errorf("viper.Unmarshal Failed: %s", err.Error())
	}
	return nil
}

func init() {
	viper.AddConfigPath("./")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../configs")
}
