package student

import (
	"reflect"

	"github.com/zopsmart/gofr/pkg/errors"
	"github.com/zopsmart/gofr/pkg/gofr"
	"github.com/zopsmart/sample-crud/models"
	"github.com/zopsmart/sample-crud/stores"
)

type student struct {
	store stores.Student
}

func New(store stores.Student) *student {
	return &student{store: store}
}

func (s *student) Find(c *gofr.Context, id string) (*models.Student, error) {
	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	student, err := s.store.Find(c, id)
	if err != nil {
		return nil, err
	}

	if reflect.DeepEqual(student, &models.Student{}) {
		return nil, errors.EntityNotFound{Entity: "student", ID: id}
	}

	return student, nil
}

func (s *student) Create(c *gofr.Context, student *models.Student) error {
	_, err := s.Find(c, student.ID)
	if err == nil {
		return errors.EntityAlreadyExists{}
	}

	return s.store.Create(c, student)
}

func (s *student) Update(c *gofr.Context, id string, student *models.Student) error {
	if id == "" {
		return errors.MissingParam{Param: []string{"id"}}
	}

	_, err := s.Find(c, id)
	if err != nil {
		return errors.EntityNotFound{Entity: "student", ID: id}
	}

	return s.store.Update(c, id, student)
}

func (s *student) Delete(c *gofr.Context, id string) error {
	if id == "" {
		return errors.MissingParam{Param: []string{"id"}}
	}

	_, err := s.Find(c, id)
	if err != nil {
		return errors.EntityNotFound{Entity: "student", ID: id}
	}

	return s.store.Delete(c, id)
}
