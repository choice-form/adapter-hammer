package model

import (
	"fmt"

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

func (ct *Adapter) SetConfig(cfg map[string]any) {
	var list []Config
	// var cfg map[string]any
	// b, _ := json.Marshal(conf)
	// json.Unmarshal(b, &cfg)
	for k, v := range cfg {
		conf := Config{
			Key:   k,
			Value: fmt.Sprint(v),
			Salt:  "",
		}
		list = append(list, conf)
	}
	ct.Configs = list
}

func (ct *Adapter) GetConfigs() []Config {
	return ct.Configs
}

func (ct *Adapter) GetConfigValue(key string) (string, error) {
	for _, cfg := range ct.Configs {
		if cfg.Key == key {
			return cfg.Value, nil
		}
	}
	return "", ErrNotFoundConfig
}
