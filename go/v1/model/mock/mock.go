package mock

import (
	m "../"
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
