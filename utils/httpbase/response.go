package httpbase

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode int, errCode interface{}, data interface{}) {
	var msg interface{}

	switch err := errCode.(type) {
	case int:
		fmt.Println(err)
		msg = GetMsg(errCode.(int))
	case []*validation.Error:
		msg = err
	default:
		fmt.Println("default")
	}
	
	g.C.JSON(
		httpCode,
		gin.H{
			"code": httpCode,
			"msg": msg,
			"data": data,
		})

	return
}