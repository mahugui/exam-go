package routers

import (
	//"net/httpbase"
	"github.com/gin-gonic/gin"

	"github.com/goexam/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode("debug")

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/products", v1.GetProducts)
		apiv1.GET("/product/:id", v1.GetProduct)
		apiv1.GET("/user_product/:user_id", v1.GetUserProducts)
		apiv1.GET("/exam_data/:user_id/:product_id", v1.ExamDataView)
	}
	return r
}