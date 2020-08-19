package model

import "github.com/fkmhrk/OpenInvoice/v1/entity"

type Trading interface {
	GetList() ([]*entity.Trading, error)
	GetListByUser(userId string) ([]*entity.Trading, error)
	GetById(id string) (*entity.Trading, error)
	Create(companyId, subject string, titleType int, workFrom, workTo, total, quotationDate, billDate, deliveryDate int64, taxRate float32, assignee, product, memo string) (*entity.Trading, error)
	Update(trading entity.Trading) (*entity.Trading, error)

	// Deletes trading and child items
	Delete(id string) error

	// Gets all trading items by trading ID
	GetItemsById(tradingId string) ([]*entity.TradingItem, error)
	CreateItem(tradingId, subject, degree, memo string, sortOrder, unitPrice int, Amount float64, taxType int) (*entity.TradingItem, error)
	// Updates specified trading Item
	UpdateItem(id, tradingId, subject, degree, memo string, sortOrder, unitPrice int, Amount float64, taxType int) (*entity.TradingItem, error)

	// Deletes specified trading Item
	SoftDeleteItem(id, tradingId string) error
}
