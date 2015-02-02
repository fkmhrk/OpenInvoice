package impl

import (
	s "../"
	m "../../model"
)

type userService struct {
	userDAO    m.UserDAO
	sessionDAO m.SessionDAO
}

const (
	MSG_ERR_NAME_EMPTY   = "name must not be empty."
	MSG_ERR_PASS_EMPTY   = "password must not be empty."
	MSG_WRONG_IDENTIFIER = "Wrong name or password."

	MSG_WRONG_TOKEN  = "Wrong token."
	MSG_SERVER_ERROR = "Server error."

	TIME_30MIN = 30 * 60
)

func NewUserSerivce(u m.UserDAO, s m.SessionDAO) *userService {
	return &userService{
		userDAO:    u,
		sessionDAO: s,
	}
}

func (s *userService) GetToken(name, pass string) s.Result {
	// input check
	if isEmpty(name) {
		return errorResult(400, MSG_ERR_NAME_EMPTY)
	}
	if isEmpty(pass) {
		return errorResult(400, MSG_ERR_PASS_EMPTY)
	}
	// get User
	user, err := s.userDAO.GetByNamePassword(name, pass)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if user == nil {
		return errorResult(400, MSG_WRONG_IDENTIFIER)
	}
	// create session
	session, err := s.sessionDAO.Create(user.Id, "read,write", TIME_30MIN)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	body := map[string]interface{}{
		"id":           user.Id,
		"access_token": session.Token,
		"token_type":   "bearer",
	}
	return jsonResult(200, body)
}
