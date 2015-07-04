package impl

import (
	s "../"
	m "../../model"
)

type userService struct {
	userDAO           m.UserDAO
	sessionDAO        m.SessionDAO
	sessionRefreshDAO m.SessionRefreshDAO
}

const (
	TIME_30MIN = 30 * 60
)

func NewUserSerivce(u m.UserDAO, s m.SessionDAO, models *m.Models) *userService {
	return &userService{
		userDAO:           u,
		sessionDAO:        s,
		sessionRefreshDAO: models.SessionRefresh,
	}
}

func (o *userService) GetToken(name, pass string) s.Result {
	// input check
	if isEmpty(name) {
		return errorResult(400, MSG_ERR_NAME_EMPTY)
	}
	if isEmpty(pass) {
		return errorResult(400, MSG_ERR_PASS_EMPTY)
	}
	// get User
	user, err := o.userDAO.GetByNamePassword(name, pass)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if user == nil {
		return errorResult(400, MSG_WRONG_IDENTIFIER)
	}
	// create session
	session, err := o.sessionDAO.Create(user.Id, string(user.Role), TIME_30MIN)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	// create refresh token
	sessionRefresh, err := o.sessionRefreshDAO.Create(user.Id, string(user.Role))
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	body := map[string]interface{}{
		"id":            user.Id,
		"access_token":  session.Token,
		"refresh_token": sessionRefresh.Token,
		"token_type":    "bearer",
	}
	return jsonResult(200, body)
}

func (o *userService) RefreshToken(token string) s.Result {
	// input check
	if isEmpty(token) {
		return errorResult(400, s.ERR_TOKEN_EMPTY)
	}
	// get refresh token
	sessionRefresh, err := o.sessionRefreshDAO.Get(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	// create access token
	session, err := o.sessionDAO.Create(sessionRefresh.UserId, sessionRefresh.Role, TIME_30MIN)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"id":           sessionRefresh.UserId,
		"access_token": session.Token,
		"token_type":   "bearer",
	}

	return jsonResult(200, body)
}

func (s *userService) GetUsers(token string) s.Result {
	// input check
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}
	// get User
	users, err := s.userDAO.GetList()
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	list := make([]interface{}, 0)
	for _, t := range users {
		list = append(list, map[string]interface{}{
			"id":           t.Id,
			"display_name": t.DisplayName,
		})
	}
	body := map[string]interface{}{
		"users": list,
	}
	return jsonResult(200, body)
}

func (s *userService) Create(token, loginName, displayName, tel, password string) s.Result {
	// get session
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}
	if !session.Role.IsAdmin() {
		return errorResult(403, MSG_NOT_AUTHORIZED)
	}
	// create
	user, err := s.userDAO.Create(loginName, displayName, "Read,Write", tel, password)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"id":           user.Id,
		"display_name": user.DisplayName,
	}
	return jsonResult(201, body)
}
