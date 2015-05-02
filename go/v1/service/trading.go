package service

type TradingService interface {
	GetListByUser(token string) Result
	Create(token, companyId, subject, product string, titleType int, workFrom, workTo, total, quotationDate, billDate int64, taxRate float32) Result
	Update(token, id, companyId, subject, product string, titleType int, workFrom, workTo, total, quotationDate, billDate int64, taxRate float32) Result

	// Gets trading items
	GetItemListByTradingId(token, tradingId string) Result
	CreateItem(token, tradingId, subject, degree, memo string, sortOrder, unitPrice, amount, taxType int) Result
	// Updates trading item
	UpdateItem(token, id, tradingId, subject, degree, memo string, sortOrder, unitPrice, amount, taxType int) Result

	// Deletes trading item
	DeleteItem(token, id, tradingId string) Result
}
