package services

import (
	"github.com/zopsmart/gofr/pkg/gofr"
	"github.com/zopsmart/sample-crud/models"
)

type Student interface {
	Find(c *gofr.Context, id string) (*models.Student, error)
	Create(c *gofr.Context, student *models.Student) error
	Update(c *gofr.Context, id string, student *models.Student) error
	Delete(c *gofr.Context, id string) error
}
