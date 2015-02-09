package mock

import (
	m "../"
)

type TradingDAO struct {
	GetListByUserResult []*m.Trading
	GetByIdResult       *m.Trading
	CreateResult        *m.Trading
	UpdateResult        *m.Trading
	GetItemsByIdResult  []*m.TradingItem
	CreateItemResult    *m.TradingItem
	UpdateItemResult    *m.TradingItem
}

func (d *TradingDAO) GetListByUser(userId string) ([]*m.Trading, error) {
	return d.GetListByUserResult, nil
}

func (d *TradingDAO) GetById(id, userId string) (*m.Trading, error) {
	return d.GetByIdResult, nil
}

func (d *TradingDAO) Create(date, companyId, subject string, titleType int, workFrom, workTo, quotationDate, billDate int64, taxRate float32, assignee, product string) (*m.Trading, error) {
	return d.CreateResult, nil
}

func (d *TradingDAO) Update(id, companyId, subject string, titleType int, workFrom, workTo, quotationDate, billDate int64, taxRate float32, assignee, product string) (*m.Trading, error) {
	return d.UpdateResult, nil
}

func (d *TradingDAO) GetItemsById(tradingId string) ([]*m.TradingItem, error) {
	return d.GetItemsByIdResult, nil
}

func (d *TradingDAO) CreateItem(tradingId, subject, degree, memo string, sortOrder, unitPrice, Amount, taxType int) (*m.TradingItem, error) {
	return d.CreateItemResult, nil
}

func (d *TradingDAO) UpdateItem(id, tradingId, subject, degree, memo string, sortOrder, unitPrice, Amount, taxType int) (*m.TradingItem, error) {
	return d.UpdateItemResult, nil
}
