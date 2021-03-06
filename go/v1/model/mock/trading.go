package mock

import (
	m "github.com/fkmhrk/OpenInvoice/v1/model"
)

type TradingDAO struct {
	GetListResult        []*m.Trading
	GetListByUserResult  []*m.Trading
	GetByIdResult        *m.Trading
	CreateResult         *m.Trading
	UpdateResult         *m.Trading
	DeleteResult         error
	GetItemsByIdResult   []*m.TradingItem
	CreateItemResult     *m.TradingItem
	UpdateItemResult     *m.TradingItem
	SoftDeleteItemResult error
}

func (d *TradingDAO) GetList() ([]*m.Trading, error) {
	return d.GetListResult, nil
}

func (d *TradingDAO) GetListByUser(userId string) ([]*m.Trading, error) {
	return d.GetListByUserResult, nil
}

func (d *TradingDAO) GetById(id string) (*m.Trading, error) {
	return d.GetByIdResult, nil
}

func (d *TradingDAO) Create(companyId, subject string, titleType int, workFrom, workTo, total, quotationDate, billDate, deliveryDate int64, taxRate float32, assignee, product, memo string) (*m.Trading, error) {
	return d.CreateResult, nil
}

func (d *TradingDAO) Update(trading m.Trading) (*m.Trading, error) {
	return d.UpdateResult, nil
}

func (d *TradingDAO) Delete(id string) error {
	return d.DeleteResult
}

func (d *TradingDAO) GetItemsById(tradingId string) ([]*m.TradingItem, error) {
	return d.GetItemsByIdResult, nil
}

func (d *TradingDAO) CreateItem(tradingId, subject, degree, memo string, sortOrder, unitPrice int, Amount float64, taxType int) (*m.TradingItem, error) {
	return d.CreateItemResult, nil
}

func (d *TradingDAO) UpdateItem(id, tradingId, subject, degree, memo string, sortOrder, unitPrice int, Amount float64, taxType int) (*m.TradingItem, error) {
	return d.UpdateItemResult, nil
}

func (d *TradingDAO) SoftDeleteItem(id, tradingId string) error {
	return d.SoftDeleteItemResult
}
