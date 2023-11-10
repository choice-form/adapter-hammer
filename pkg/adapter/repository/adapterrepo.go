package repository

import (
	"context"
	"errors"

	"github.com/choice-form/adapter-hammer/pkg/adapter/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AdapterRepo struct {
	db *gorm.DB
}

func NewAdapterRepo(db *gorm.DB) *AdapterRepo {
	return &AdapterRepo{
		db: db,
	}
}

func (ct *AdapterRepo) Create(c context.Context, connectorID string, conf map[string]any) error {
	connector := &model.Adapter{}
	db := ct.db.WithContext(c)
	tx := db.Model(&model.Adapter{}).Where("connector_id = ?", connectorID).First(connector)
	// 记录不存在，创建新的连机器适配器账号
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		newConnector := &model.Adapter{
			ConnectorID: connectorID,
		}
		newConnector.SetConfig(conf)
		if _tx := db.Create(newConnector); _tx.Error != nil {
			return _tx.Error
		}
	} else {
		return tx.Error
	}
	return nil
}

func (ct *AdapterRepo) FindByConnectorID(c context.Context, connectorID string) (*model.Adapter, error) {
	connector := new(model.Adapter)
	db := ct.db.WithContext(c)
	tx := db.Preload(clause.Associations).First(connector)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return connector, nil
}

func (ct *AdapterRepo) Update(c context.Context, connectorID string, confs []model.Config) error {
	var connector *model.Adapter
	db := ct.db.WithContext(c)
	tx := db.Where("connector_id = ?", connectorID).Joins("config").First(connector)
	if tx.Error != nil {
		return tx.Error
	}
	connector.Configs = configMerge(connector.Configs, confs)
	return db.Save(connector).Error
}

func (ct *AdapterRepo) Delete(c context.Context, connectID string) error {
	return ct.db.WithContext(c).Where("connector_id = ?", connectID).Delete(&model.Adapter{}).Error
}

func configMerge(origin, replace []model.Config) []model.Config {
	for k, v := range origin {
		for _, val := range replace {
			if v.ID == val.ID {
				origin[k] = val
			}
		}
	}
	return origin
}
