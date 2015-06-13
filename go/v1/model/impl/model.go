package impl

import (
	m "../"
)

func NewModels(connection *Connection) *m.Models {
	logger := NewLogger()
	return &m.Models{
		User:    NewUserDAO(connection),
		Session: NewSessionDAO(connection),
		Company: NewCompanyDAO(connection),
		Trading: NewTradingDAO(connection, logger),
		Env:     NewEnvDAO(connection),
		Seq:     NewSeqDAO(connection),
		Logger:  logger,
	}
}
