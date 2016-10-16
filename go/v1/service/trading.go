package service

import (
	m "../model"
)

type TradingService interface {
	GetListByUser() Result
	Create(token, companyId, subject, product, memo string, titleType int, workFrom, workTo, total, quotationDate, billDate, deliveryDate int64, taxRate float32) Result
	Update(token string, trading Trading) Result

	// Deletes trading and child items
	Delete(token, tradingId string) Result

	// Gets trading items
	GetItemListByTradingId(tradingId string) Result
	CreateItem(token, tradingId, subject, degree, memo string, sortOrder, unitPrice int, amount float64, taxType int) Result
	// Updates trading item
	UpdateItem(token, id, tradingId, subject, degree, memo string, sortOrder, unitPrice int, amount float64, taxType int) Result

	// Deletes trading item
	DeleteItem(token, id, tradingId string) Result

	// Gets next number
	GetNextNumber(token, seqType string, date int64) Result
}

type Trading struct {
	m.Trading
}
