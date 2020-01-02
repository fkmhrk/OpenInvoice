package response

import "github.com/mokelab-go/server/entity"

// Error response
func Error(status int, msg string) entity.Response {
	return entity.Response{
		Status: status,
		Body: map[string]interface{}{
			"msg": msg,
		},
	}
}

const (
	MSG_ERR_NAME_EMPTY       = "name must not be empty."
	MSG_ERR_PASS_EMPTY       = "password must not be empty."
	MSG_ERR_ID_EMPTY         = "id must not be empty."
	MSG_ERR_DATE_EMPTY       = "date must not be empty."
	MSG_ERR_COMPANY_ID_EMPTY = "company_id must not be empty."
	MSG_ERR_SUBJECT_EMPTY    = "subject must not be empty."
	MSG_ERR_PRODUCT_EMPTY    = "product must not be empty."
	MSG_WRONG_IDENTIFIER     = "Wrong name or password."
	MSG_TRADING_NOT_FOUND    = "Trading not found."

	MSG_WRONG_TOKEN    = "Wrong token."
	MSG_NOT_AUTHORIZED = "User not authorized."
	MSG_SERVER_ERROR   = "Server error."

	MSG_INVALID_SEQUENCE_TYPE = "Sequence type must be quotation/delivery/bill."
	MSG_TOKEN_EMPTY           = "token must not be empty."
)
