package model

type ConfigType string

var (
	TypeAuth      ConfigType = "auth"
	TypeCallBacks ConfigType = "callbacks"
	TypeWebhooks  ConfigType = "webhooks"
)

type Config struct {
	Common
	AdapterID uint       `json:"Adapter_id" gorm:"index"`
	Key       string     `json:"key" gorm:"index"`
	Value     string     `json:"value,omitempty"`
	Type      ConfigType `json:"type,omitempty"`
	Salt      string     `json:"-"`
}

func (c *Config) TableName() string {
	return "config"
}
