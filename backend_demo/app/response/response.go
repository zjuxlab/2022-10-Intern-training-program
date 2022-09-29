package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SendResponse(c echo.Context, code int, msg string, data ...interface{}) error{
	return c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
