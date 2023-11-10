package repository

import (
	"fmt"

	"github.com/choice-form/adapter-hammer/pkg/adapter/model"
	"github.com/choice-form/adapter-hammer/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

var client *PGClient

// func init() {
// 	cfg := config.GetConfig()
// 	pg := NewDBClient(&cfg.Postgres)
// 	pgClient = pg
// }

type PGClient struct {
	config *Config
	DB     *gorm.DB
}

func InitClient(cfg *Config) {
	client = NewClient(cfg)
}

func GetClient() *PGClient {
	if client == nil {
		logger.Error("not found db client or not init db client")
	}
	return client
}

func NewClient(cfg *Config) *PGClient {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Migrate(db)

	return &PGClient{
		config: &Config{},
		DB:     db,
	}
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Adapter{},
		&model.Config{},
	)
}
