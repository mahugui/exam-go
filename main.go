package main

import (
	//"fmt"
	"time"
	"net/http"

	"github.com/goexam/conf"
	"github.com/goexam/models"
	"github.com/goexam/routers"
	"github.com/goexam/cmodels"
)

func main()  {
	conf.Setup()
	models.Setup()
	cmodels.Setup()

	routersInit := routers.InitRouter()

	s := &http.Server{
		Addr: ":8080",
		Handler: routersInit,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
