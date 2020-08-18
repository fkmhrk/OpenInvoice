package mock

import (
	"github.com/fkmhrk/OpenInvoice/v1/entity"
)

type TradingDAO struct {
	GetListResult        []*entity.Trading
	GetListByUserResult  []*entity.Trading
	GetByIdResult        *entity.Trading
	CreateResult         *entity.Trading
	UpdateResult         *entity.Trading
	DeleteResult         error
	GetItemsByIdResult   []*entity.TradingItem
	CreateItemResult     *entity.TradingItem
	UpdateItemResult     *entity.TradingItem
	SoftDeleteItemResult error
}

func (d *TradingDAO) GetList() ([]*entity.Trading, error) {
	return d.GetListResult, nil
}

func (d *TradingDAO) GetListByUser(userId string) ([]*entity.Trading, error) {
	return d.GetListByUserResult, nil
}

func (d *TradingDAO) GetById(id string) (*entity.Trading, error) {
	return d.GetByIdResult, nil
}

func (d *TradingDAO) Create(companyId, subject string, titleType int, workFrom, workTo, total, quotationDate, billDate, deliveryDate int64, taxRate float32, assignee, product, memo string) (*entity.Trading, error) {
	return d.CreateResult, nil
}

func (d *TradingDAO) Update(trading entity.Trading) (*entity.Trading, error) {
	return d.UpdateResult, nil
}

func (d *TradingDAO) Delete(id string) error {
	return d.DeleteResult
}

func (d *TradingDAO) GetItemsById(tradingId string) ([]*entity.TradingItem, error) {
	return d.GetItemsByIdResult, nil
}

func (d *TradingDAO) CreateItem(tradingId, subject, degree, memo string, sortOrder, unitPrice int, Amount float64, taxType int) (*entity.TradingItem, error) {
	return d.CreateItemResult, nil
}

func (d *TradingDAO) UpdateItem(id, tradingId, subject, degree, memo string, sortOrder, unitPrice int, Amount float64, taxType int) (*entity.TradingItem, error) {
	return d.UpdateItemResult, nil
}

func (d *TradingDAO) SoftDeleteItem(id, tradingId string) error {
	return d.SoftDeleteItemResult
}
