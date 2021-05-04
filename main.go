package main

import (
	"github.com/zopsmart/gofr/pkg/gofr"
	"github.com/zopsmart/sample-crud/handlers"
	"github.com/zopsmart/sample-crud/stores"
)

func main() {
	store := stores.New()
	handler := handlers.New(store)

	k := gofr.New()

	k.GET("/student/{id:[0-9]+}", handler.Find)
	k.POST("/student", handler.Create)
	k.PUT("/student/{id:[0-9]+}", handler.Update)
	k.DELETE("/student/{id:[0-9]+}", handler.Delete)

	k.Start()
}
