package base

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type baseReponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Error  interface{} `json:"error"`
	Data   interface{} `json:"data"`
}

func SetResponseSuccess(data interface{}) *baseReponse {
	br := baseReponse{}
	br.Code = http.StatusOK
	br.Status = "ok"
	br.Data = data

	return &br
}

func SetResponseError(err interface{}) *baseReponse {
	br := baseReponse{}
	br.Code = http.StatusBadRequest
	br.Status = "bad_request"
	br.Data = err

	return &br
}

func (br *baseReponse) Thrown(c echo.Context) error {
	return c.JSON(br.Code, br)
}
