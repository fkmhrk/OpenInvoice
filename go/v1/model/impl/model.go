package impl

import (
	m "github.com/fkmhrk/OpenInvoice/v1/model"
)

func NewModels(connection *Connection) *m.Models {
	logger := NewLogger()
	return &m.Models{
		User:           NewUserDAO(connection),
		Session:        NewSessionDAO(connection),
		SessionRefresh: NewSessionRefreshDAO(connection),
		Company:        NewCompanyDAO(connection),
		Trading:        NewTradingDAO(connection, logger),
		Env:            NewEnvDAO(connection),
		Seq:            NewSeqDAO(connection),
		Logger:         logger,
	}
}
