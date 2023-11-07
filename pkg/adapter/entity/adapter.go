package entity

import (
	"fmt"

	"github.com/choice-form/adapter-hammer/pkg/utils"
	"github.com/choice-form/dingtalk-adapter/pkg/dingding/certificate"
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

func (ct *Adapter) GetConfig() *certificate.Config {
	var cfg = new(certificate.Config)
	// var conf map[string]any
	conf := make(map[string]any)
	for _, v := range ct.Configs {
		conf[v.Key] = v.Value
	}
	utils.Map2Struct(conf, cfg)
	return cfg
}
