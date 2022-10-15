package exception

import (
	"encoding/json"
)

const (
	CodeSuccess = iota
	CodeErrRequestNotValid
	CodeErrQueryDB
	CodeErrDataNotFound
	CodeInternalServerError
	CodeCreated
	CodeBadRequest
)

var CodeMapping = map[int]string{
	CodeSuccess:             "Success",
	CodeErrRequestNotValid:  "Request not valid",
	CodeErrQueryDB:          "There is error while query DB",
	CodeErrDataNotFound:     "No Found",
	CodeInternalServerError: "Internal server error",
	CodeCreated:             "Success created",
	CodeBadRequest:          "Bad Request",
}

var HttpCodeMapping = map[int]int{
	CodeSuccess:             200,
	CodeErrRequestNotValid:  400,
	CodeErrQueryDB:          500,
	CodeErrDataNotFound:     404,
	CodeInternalServerError: 500,
	CodeCreated:             201,
	CodeBadRequest:          400,
}

type ErrorWithCode struct {
	Code            int
	HttpResponeCode int
	Msg             string
}

func (c ErrorWithCode) Error() string {
	b, _ := json.Marshal(&c)
	return string(b)
}

func NewErrorMsg(code int, err error) error {
	msg := CodeMapping[code]
	httpCode := HttpCodeMapping[code]
	return ErrorWithCode{
		HttpResponeCode: httpCode,
		Code:            code,
		Msg:             msg,
	}
}
