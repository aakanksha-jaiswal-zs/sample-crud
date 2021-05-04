package handlers

import (
	"github.com/zopsmart/gofr/pkg/gofr"
	"github.com/zopsmart/sample-crud/models"
	"github.com/zopsmart/sample-crud/stores"
)

type student struct {
	store stores.Student
}

func New(store stores.Student) student {
	return student{store: store}
}

func (s student) Find(c *gofr.Context) (interface{}, error) {
	id := c.PathParam("id")
	return s.store.Find(id)
}

func (s student) Create(c *gofr.Context) (interface{}, error) {
	var student models.Student

	if err := c.Bind(&student); err != nil {
		return nil, err
	}

	return nil, s.store.Create(student)
}

func (s student) Update(c *gofr.Context) (interface{}, error) {
	id := c.PathParam("id")

	var student models.Student

	if err := c.Bind(&student); err != nil {
		return nil, err
	}

	return nil, s.store.Update(id, student)
}

func (s student) Delete(c *gofr.Context) (interface{}, error) {
	id := c.PathParam("id")

	return nil, s.store.Delete(id)
}
