package db

import (
	"context"
	"errors"

	"github.com/choice-form/adapter-hammer/pkg/adapter/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AdapterRepo struct {
}

func (ct *AdapterRepo) Create(c context.Context, connectorID string, conf map[string]any) error {
	connector := &entity.Adapter{}
	db := GetClient().DB.WithContext(c)
	tx := db.Model(&entity.Adapter{}).Where("connector_id = ?", connectorID).First(connector)
	// 记录不存在，创建新的连机器适配器账号
	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		newConnector := &entity.Adapter{
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

func (ct *AdapterRepo) FindByConnectorID(c context.Context, connectorID string) (*entity.Adapter, error) {
	connector := new(entity.Adapter)
	db := GetClient().DB.WithContext(c)
	tx := db.Preload(clause.Associations).First(connector)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return connector, nil
}

func (ct *AdapterRepo) Update(c context.Context, connectorID string, confs []entity.Config) error {
	var connector *entity.Adapter
	db := GetClient().DB.WithContext(c)
	tx := db.Where("connector_id = ?", connectorID).Joins("config").First(connector)
	if tx.Error != nil {
		return tx.Error
	}
	connector.Configs = configMerge(connector.Configs, confs)
	return db.Save(connector).Error
}

func (ct *AdapterRepo) Delete(c context.Context, connectID string) error {
	return GetClient().DB.WithContext(c).Where("connector_id = ?", connectID).Delete(&entity.Adapter{}).Error
}

func configMerge(origin, replace []entity.Config) []entity.Config {
	for k, v := range origin {
		for _, val := range replace {
			if v.ID == val.ID {
				origin[k] = val
			}
		}
	}
	return origin
}
