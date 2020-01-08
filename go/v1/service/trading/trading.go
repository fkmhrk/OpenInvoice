package trading

import (
	"github.com/fkmhrk/OpenInvoice/v1/model/session"
	"github.com/fkmhrk/OpenInvoice/v1/model/trading"
	"github.com/mokelab-go/server/entity"
)

type Service interface {
	GetListByUser() entity.Response
	GetTradingByID(id string) entity.Response
	Create(session *session.Session, companyId, subject, product, memo string, titleType int, workFrom, workTo, total, quotationDate, billDate, deliveryDate int64, taxRate float32) entity.Response
	Update(trading Trading) entity.Response

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

type Trading struct {
	trading.Trading
}
