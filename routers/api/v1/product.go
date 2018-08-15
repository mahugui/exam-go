package v1

import (
	"fmt"
	"net/http"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/goexam/models"
	"github.com/goexam/utils/httpbase"
)

func GetProducts(c *gin.Context)  {
	products, err := models.GetProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, "失败")
	}
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context)  {
	appG := httpbase.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors(){
		fmt.Println(valid.Errors)
		appG.Response(http.StatusOK, httpbase.INVALID_PARAMS, nil)
		return
	}

	exists, err := models.ExistProductByID(id)
	fmt.Println(exists)
	if err != nil{
		appG.Response(http.StatusOK, httpbase.ERROR, nil)
		return
	}
	if !exists{
		appG.Response(http.StatusOK, httpbase.INVALID_PARAMS, nil)
		return
	}

	product, err := models.GetProduct(id)
	if err != nil{
		appG.Response(http.StatusOK, httpbase.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, httpbase.SUCCESS, product)
}