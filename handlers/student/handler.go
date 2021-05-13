package student

import (
	"github.com/zopsmart/gofr/pkg/gofr"
	"github.com/zopsmart/sample-crud/models"
	"github.com/zopsmart/sample-crud/services"
)

type student struct {
	service services.Student
}

func New(store services.Student) *student {
	return &student{service: store}
}

func (s *student) Find(c *gofr.Context) (interface{}, error) {
	id := c.PathParam("id")
	return s.service.Find(c, id)
}

func (s *student) Create(c *gofr.Context) (interface{}, error) {
	var student models.Student

	if err := c.Bind(&student); err != nil {
		return nil, err
	}

	return nil, s.service.Create(c, &student)
}

func (s *student) Update(c *gofr.Context) (interface{}, error) {
	id := c.PathParam("id")

	var student models.Student

	if err := c.Bind(&student); err != nil {
		return nil, err
	}

	return nil, s.service.Update(c, id, &student)
}

func (s *student) Delete(c *gofr.Context) (interface{}, error) {
	id := c.PathParam("id")

	return nil, s.service.Delete(c, id)
}
