package base

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseReponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Error  interface{} `json:"error"`
	Data   interface{} `json:"data"`
}

func SetResponseSuccess(data interface{}) (br *BaseReponse) {
	br.Code = http.StatusOK
	br.Status = "ok"
	br.Data = data

	return
}

func SetResponseError(err interface{}) (br *BaseReponse) {
	br.Code = http.StatusBadRequest
	br.Status = "bad_request"
	br.Data = err

	return
}

func (br *BaseReponse) Thrown(c echo.Context) error {
	return c.JSON(br.Code, br)
}
