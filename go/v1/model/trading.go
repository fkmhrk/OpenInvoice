package model

type TradingDAO interface {
	GetListByUser(userId string) ([]*Trading, error)
	Create(date, companyId, subject string, workFrom, workTo int64, assignee, product string) (*Trading, error)

	// Gets all trading items by trading ID
	GetItemsById(tradingId string) ([]*TradingItem, error)
	CreateItem(tradingId, subject string, unitPrice, Amount int,
		degree string, taxType int, memo string) (*TradingItem, error)
}

type Trading struct {
	Id           string
	CompanyId    string
	Subject      string
	WorkFrom     int64
	WorkTo       int64
	AssigneeId   string
	Product      string
	CreatedTime  int64
	ModifiedTime int64
}

type TradingItem struct {
	Id        string
	TradingId string
	SortOrder int
	Subject   string
	UnitPrice int
	Amount    int
	Degree    string
	TaxType   int
	Memo      string
}