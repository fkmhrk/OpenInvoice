package service

type TradingService interface {
	GetListByUser(token string) Result
	Create(token, date, companyId, subject, product string, workFrom, workTo int64) Result
	Update(token, id, companyId, subject, product string, titleType int, workFrom, workTo int64) Result

	// Gets trading items
	GetItemListByTradingId(token, tradingId string) Result
	CreateItem(token, tradingId, subject, degree, memo string, sortOrder, unitPrice, amount, taxType int) Result
}
