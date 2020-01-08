package mock

import (
	m "github.com/fkmhrk/OpenInvoice/v1/model"
)

func NewMock() *m.Models {
	return &m.Models{
		User:           &UserDAO{},
		Session:        &SessionDAO{},
		SessionRefresh: &SessionRefreshDAO{},
		Company:        &CompanyDAO{},
		Trading:        &TradingDAO{},
		Env:            &EnvDAO{},
		Seq:            &SeqDAO{},
	}
}
