package helper

import (
	"fmt"

	"github.com/labstack/echo"
)

type ResponseData struct {
	Status int
	Data   interface{}
}

func RespondSuccess(w echo.Context, status int, payload interface{}) {
	fmt.Println("status ", status)
	var res ResponseData

	res.Status = status
	res.Data = payload

	w.JSON(status, res)
}
