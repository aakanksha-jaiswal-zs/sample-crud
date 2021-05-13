package main

import (
	"github.com/zopsmart/gofr/pkg/gofr"
	student2 "github.com/zopsmart/sample-crud/handlers/student"
	student3 "github.com/zopsmart/sample-crud/services/student"
	"github.com/zopsmart/sample-crud/stores/student"
)

func main() {
	store := student.New()
	service := student3.New(store)
	handler := student2.New(service)

	k := gofr.New()

	k.GET("/student/{id:[0-9]+}", handler.Find)
	k.POST("/student", handler.Create)
	k.PUT("/student/{id:[0-9]+}", handler.Update)
	k.DELETE("/student/{id:[0-9]+}", handler.Delete)

	k.Start()
}
