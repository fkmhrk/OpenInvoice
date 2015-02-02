package service

type TradingService interface {
	GetListByUser(token string) Result
}
