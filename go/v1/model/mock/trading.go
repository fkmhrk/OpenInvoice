package mock

import (
	m "../"
)

type TradingDAO struct {
	GetListByUserResult []*m.Trading
}

func (d *TradingDAO) GetListByUser(userId string) ([]*m.Trading, error) {
	return d.GetListByUserResult, nil
}
