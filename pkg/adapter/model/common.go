package model

import (
	"time"

	"gorm.io/gorm"
)

type Common struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `json:"-" gorm:"index"`
	UpdateAt  time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
