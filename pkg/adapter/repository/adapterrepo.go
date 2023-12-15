package repository

import (
	"context"

	"github.com/choice-form/adapter-hammer/pkg/adapter/model"
	"github.com/pkg/errors"
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

func (ct *AdapterRepo) Create(c context.Context, connectorID string, confs []model.Config) (*model.Adapter, error) {
	connector := &model.Adapter{}
	db := ct.db.WithContext(c)
	tx := db.Model(&model.Adapter{}).Where("connector_id = ?", connectorID).First(connector)
	// tx 出错了
	if tx.Error != nil {
		// 记录不存在，创建新的连机器适配器账号
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			newConnector := &model.Adapter{
				ConnectorID: connectorID,
			}
			newConnector.SetConfig(confs)
			_tx := db.Create(newConnector)
			if _tx.Error != nil {
				return nil, _tx.Error
			}
			return newConnector, nil
		} else {
			// 其他错误直接返回错误
			return nil, tx.Error
		}
	}

	return nil, errors.New("adapter already exists")
}

func (ct *AdapterRepo) FindByConnectorID(c context.Context, connectorID string) (*model.Adapter, error) {
	connector := new(model.Adapter)
	db := ct.db.WithContext(c)
	tx := db.Preload(clause.Associations).Where("connector_id = ?", connectorID).First(connector)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return connector, nil
}

func (ct *AdapterRepo) Update(c context.Context, connectorID string, adap model.Adapter) error {
	connector := new(model.Adapter)
	db := ct.db.WithContext(c)
	tx := db.Preload(clause.Associations).Where("connector_id = ?", connectorID).First(connector)
	if tx.Error != nil {
		return tx.Error
	}

	for i := 0; i < len(adap.Configs); i++ {
		adap.Configs[i].AdapterID = connector.ID
	}

	connector.Configs = adap.Configs
	return db.Save(connector).Error
}

func (ct *AdapterRepo) Delete(c context.Context, connectID string) error {
	return ct.db.WithContext(c).Where("connector_id = ?", connectID).Delete(&model.Adapter{}).Error
}

func (ct *AdapterRepo) GetAllConnector(c context.Context) ([]model.Adapter, error) {
	list := make([]model.Adapter, 0)
	tx := ct.db.WithContext(c).Preload(clause.Associations).Find(&list)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return list, nil
}
