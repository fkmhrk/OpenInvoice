package entity

// Trading is entity
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
	DeliveryDate    int64
	DeliveryNumber  string
	TaxRate         float32
	AssigneeId      string
	Product         string
	Memo            string
	CreatedTime     int64
	ModifiedTime    int64
}

// TradingItem is entity
type TradingItem struct {
	Id        string
	TradingId string
	SortOrder int
	Subject   string
	UnitPrice int
	Amount    float64
	Degree    string
	TaxType   int
	Memo      string
}
