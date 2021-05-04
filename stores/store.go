package stores

import (
	"github.com/zopsmart/gofr/pkg/errors"
	"github.com/zopsmart/sample-crud/db"
	"github.com/zopsmart/sample-crud/models"
)

type student struct {
	db []models.Student
}

func New() student {
	return student{db: db.Students}
}

func (s student) Find(id string) (models.Student, error) {
	for _, student := range s.db {
		if student.ID == id {
			return student, nil
		}
	}

	return models.Student{}, errors.EntityNotFound{Entity: "student", ID: id}
}

func (s student) Create(student models.Student) error {
	_, err := s.Find(student.ID)
	if err == nil {
		return errors.EntityAlreadyExists{}
	}
	s.db = append(s.db, student)

	return nil
}

func (s student) Update(id string, student models.Student) error {
	for i := range s.db {
		if s.db[i].ID == id {
			s.db[i].GPA = student.GPA
			s.db[i].Name = student.Name
			s.db[i].Email = student.Email
			return nil
		}
	}

	return errors.EntityNotFound{Entity: "student", ID: id}
}

func (s student) Delete(id string) error {
	for i, student := range s.db {
		if student.ID == id {
			s.db = append(s.db[:i], s.db[i+1:]...)
			return nil
		}
	}

	return errors.EntityNotFound{Entity: "student", ID: id}
}
