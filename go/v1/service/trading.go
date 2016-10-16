package service

import (
	m "../model"
)

type TradingService interface {
	GetListByUser() Result
	Create(session *m.Session, companyId, subject, product, memo string, titleType int, workFrom, workTo, total, quotationDate, billDate, deliveryDate int64, taxRate float32) Result
	Update(trading Trading) Result

	// Deletes trading and child items
	Delete(tradingId string) Result

	// Gets trading items
	GetItemListByTradingId(tradingId string) Result
	CreateItem(tradingId, subject, degree, memo string, sortOrder, unitPrice int, amount float64, taxType int) Result
	// Updates trading item
	UpdateItem(id, tradingId, subject, degree, memo string, sortOrder, unitPrice int, amount float64, taxType int) Result

	// Deletes trading item
	DeleteItem(id, tradingId string) Result

	// Gets next number
	GetNextNumber(seqType string, date int64) Result
}

type Trading struct {
	m.Trading
}
