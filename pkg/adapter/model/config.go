package model

type ConfigType string

var (
	TypeAuth     ConfigType = "auth"
	TypeWebhooks ConfigType = "webhooks"
	TypeOther    ConfigType = "other"
	TypeBase     ConfigType = "base"
)

type Config struct {
	Common
	AdapterID uint       `json:"Adapter_id" gorm:"index"`
	Key       string     `json:"key" gorm:"index"`
	Value     *string    `json:"value,omitempty"`
	Type      ConfigType `json:"type,omitempty"`
	Encrypt   *bool      `json:"encrypt"`
}

func (c *Config) TableName() string {
	return "config"
}
