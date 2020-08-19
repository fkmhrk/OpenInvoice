package mock

import (
	"github.com/fkmhrk/OpenInvoice/v1/model"
)

func NewMock() *model.Models {
	return &model.Models{
		User:           &UserDAO{},
		Session:        &SessionDAO{},
		SessionRefresh: &SessionRefreshDAO{},
		Company:        &CompanyDAO{},
		Trading:        &TradingDAO{},
		Env:            &EnvDAO{},
		Seq:            &SeqDAO{},
	}
}
