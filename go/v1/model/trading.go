package model

type TradingDAO interface {
	GetListByUser(userId string) ([]*Trading, error)
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
