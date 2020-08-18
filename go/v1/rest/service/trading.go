package service

import (
	e "github.com/fkmhrk/OpenInvoice/v1/entity"
	"github.com/mokelab-go/server/entity"
)

type Trading interface {
	GetListByUser() entity.Response
	GetTradingByID(id string) entity.Response
	Create(session *e.Session, companyId, subject, product, memo string, titleType int, workFrom, workTo, total, quotationDate, billDate, deliveryDate int64, taxRate float32) entity.Response

	Update(trading TradingData) entity.Response

	// Deletes trading and child items
	Delete(tradingId string) entity.Response

	// Gets trading items
	GetItemListByTradingId(tradingId string) entity.Response
	CreateItem(tradingId, subject, degree, memo string, sortOrder, unitPrice int, amount float64, taxType int) entity.Response
	// Updates trading item
	UpdateItem(id, tradingId, subject, degree, memo string, sortOrder, unitPrice int, amount float64, taxType int) entity.Response

	// Deletes trading item
	DeleteItem(id, tradingId string) entity.Response

	// Gets next number
	GetNextNumber(seqType string, date int64) entity.Response
}

type TradingData struct {
	e.Trading
}
