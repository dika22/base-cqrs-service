package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Result interface{} `json:"result"`
	Message string `json:"message"`
	Status bool `json:"status"`
	StatusCode int64 `json:"status_code"`
}


func JSONSuccess(c echo.Context, result interface{}, msg string) error {
	return c.JSON(http.StatusOK, Response{
		Result:     result,
		Message:    msg,
		Status:     true,
		StatusCode: http.StatusOK,
	})
 }
 

 func JSONResponse(c echo.Context, code int, status bool, message string, data interface{}) error {
    return c.JSON(code, map[string]interface{}{
        "status":  status,
        "message": message,
        "data":    data,
    })
}