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
		"is_admin":      user.Role.IsAdmin(),
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
	session, err := o.sessionDAO.Create(sessionRefresh.UserId, string(sessionRefresh.Role), TIME_30MIN)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"id":           sessionRefresh.UserId,
		"access_token": session.Token,
		"is_admin":     sessionRefresh.Role.IsAdmin(),
		"token_type":   "bearer",
	}

	return jsonResult(200, body)
}

func (s *userService) GetUsers() s.Result {
	// get User
	users, err := s.userDAO.GetList()
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	list := make([]interface{}, 0)
	for _, t := range users {
		list = append(list, map[string]interface{}{
			"id":           t.Id,
			"login_name":   t.LoginName,
			"display_name": t.DisplayName,
			"tel":          t.Tel,
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
		"login_name":   user.LoginName,
		"display_name": user.DisplayName,
		"tel":          user.Tel,
	}
	return jsonResult(201, body)
}

func (s *userService) Update(token, id, loginName, displayName, tel, password string) s.Result {
	// get session
	session, err := s.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}
	if !session.Role.IsAdmin() {
		if session.UserId != id {
			return errorResult(403, MSG_NOT_AUTHORIZED)
		}
	}
	// update
	_, err = s.userDAO.Update(id, loginName, displayName, "", tel, password)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}

	body := map[string]interface{}{
		"id":           id,
		"login_name":   loginName,
		"display_name": displayName,
		"tel":          tel,
	}
	return jsonResult(200, body)
}

func (o *userService) Delete(token, id string) s.Result {
	// get session
	session, err := o.sessionDAO.GetByToken(token)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if session == nil {
		return errorResult(401, MSG_WRONG_TOKEN)
	}
	if !session.Role.IsAdmin() {
		return errorResult(403, MSG_NOT_AUTHORIZED)
	}
	// get user
	user, err := o.userDAO.GetById(id)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	if user.Role.IsAdmin() {
		return errorResult(403, MSG_NOT_AUTHORIZED)
	}
	// delete
	err = o.userDAO.Delete(id)
	if err != nil {
		return errorResult(500, MSG_SERVER_ERROR)
	}
	return &result{
		status: 204,
		body:   "",
	}
}
