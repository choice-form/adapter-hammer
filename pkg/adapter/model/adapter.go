package model

import (
	"github.com/pkg/errors"
)

var (
	ErrNotFoundConfig = errors.New("not found config")
)

type Adapter struct {
	Common
	ConnectorID string   `json:"connector_id" gorm:"column:connector_id;uniqueIndex"`
	Configs     []Config `json:"configs"`
}

func (ct *Adapter) TableName() string {
	return "adapter"
}

func (ct *Adapter) SetConfig(confs []Config) {
	ct.Configs = confs
}

func (ct *Adapter) GetConfigs() []Config {
	return ct.Configs
}

func (ct *Adapter) GetConfig(key string) (*Config, error) {
	for i := 0; i < len(ct.Configs); i++ {
		if ct.Configs[i].Key == key {
			return &ct.Configs[i], nil
		}
	}
	return &Config{}, ErrNotFoundConfig
}

func (ct *Adapter) GetConfigValue(key string) *string {
	for _, cfg := range ct.Configs {
		if cfg.Key == key {
			return cfg.Value
		}
	}
	return nil
}
