package mock

import (
	m "../"
)

type TradingDAO struct {
	GetListByUserResult []*m.Trading
	CreateResult        *m.Trading
	GetItemsByIdResult  []*m.TradingItem
	CreateItemResult    *m.TradingItem
}

func (d *TradingDAO) GetListByUser(userId string) ([]*m.Trading, error) {
	return d.GetListByUserResult, nil
}

func (d *TradingDAO) Create(date, companyId, subject string, workFrom, workTo int64, assignee, product string) (*m.Trading, error) {
	return d.CreateResult, nil
}

func (d *TradingDAO) GetItemsById(tradingId string) ([]*m.TradingItem, error) {
	return d.GetItemsByIdResult, nil
}

func (d *TradingDAO) CreateItem(tradingId, subject, degree, memo string, sortOrder, unitPrice, Amount, taxType int) (*m.TradingItem, error) {
	return d.CreateItemResult, nil
}
