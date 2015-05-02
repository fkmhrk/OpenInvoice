package model

type TradingDAO interface {
	GetListByUser(userId string) ([]*Trading, error)
	GetById(id, userId string) (*Trading, error)
	Create(companyId, subject string, titleType int, workFrom, workTo, total, quotationDate, billDate int64, taxRate float32, assignee, product string) (*Trading, error)
	Update(id, companyId, subject string, titleType int, workFrom, workTo, quotationDate, billDate int64, taxRate float32, assignee, product string) (*Trading, error)

	// Gets all trading items by trading ID
	GetItemsById(tradingId string) ([]*TradingItem, error)
	CreateItem(tradingId, subject, degree, memo string, sortOrder, unitPrice, Amount, taxType int) (*TradingItem, error)
	// Updates specified trading Item
	UpdateItem(id, tradingId, subject, degree, memo string, sortOrder, unitPrice, Amount, taxType int) (*TradingItem, error)

	// Deletes specified trading Item
	SoftDeleteItem(id, tradingId string) error
}

type Trading struct {
	Id              string
	CompanyId       string
	Subject         string
	TitleType       int
	WorkFrom        int64
	WorkTo          int64
	Total           int64
	QuotationDate   int64
	QuotationNumber string
	BillDate        int64
	BillNumber      string
	TaxRate         float32
	AssigneeId      string
	Product         string
	CreatedTime     int64
	ModifiedTime    int64
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
