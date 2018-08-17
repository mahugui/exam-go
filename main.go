package main

import (
	"fmt"
	"os"
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

        port := os.Getenv("PORT")
        fmt.Println(port)
	s := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: routersInit,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
