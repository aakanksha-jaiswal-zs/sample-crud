package stores

import "github.com/zopsmart/sample-crud/models"

type Student interface {
	Find(id string) (models.Student, error)
	Create(student models.Student) error
	Update(id string, student models.Student) error
	Delete(id string) error
}
