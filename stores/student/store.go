package student

import (
	"strings"

	"github.com/zopsmart/gofr/pkg/errors"
	"github.com/zopsmart/gofr/pkg/gofr"
	"github.com/zopsmart/sample-crud/models"
)

type student struct{}

func New() *student {
	return &student{}
}

func (s *student) Find(c *gofr.Context, id string) (*models.Student, error) {
	session := c.Cassandra.Session

	query := `SELECT id, name, email, gpa FROM students where id = ?`

	scanner := session.Query(query, id).WithContext(c.Context).Iter().Scanner()

	var student models.Student

	for scanner.Next() {
		err := scanner.Scan(&student.ID, &student.Name, &student.Email, &student.GPA)
		if err != nil {
			return nil, errors.DB{Err: err}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.DB{Err: err}
	}

	return &student, nil
}

func (s *student) Create(c *gofr.Context, student *models.Student) error {
	session := c.Cassandra.Session

	query := `INSERT INTO students( id, name, email, gpa) VALUES (?, ?, ?, ?)`

	err := session.Query(query, student.ID, student.Name, student.Email, student.GPA).WithContext(c.Context).Exec()
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

func (s *student) Update(c *gofr.Context, id string, student *models.Student) error {
	session := c.Cassandra.Session

	query, values := generateSetClause(student)
	values = append(values, id)

	err := session.Query(query, values...).WithContext(c.Context).Exec()
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

func generateSetClause(student *models.Student) (query string, values []interface{}) {
	query = `UPDATE students SET`

	if student.Name != "" {
		query += ` name = ?,`
		values = append(values, student.Name)
	}

	if student.GPA != 0 {
		query += ` gpa = ?,`
		values = append(values, student.GPA)
	}

	if student.Email != "" {
		query += ` email = ?,`
		values = append(values, student.Email)
	}

	query = strings.TrimSuffix(query, ",")
	query += ` WHERE id = ?`

	return query, values
}

func (s *student) Delete(c *gofr.Context, id string) error {
	session := c.Cassandra.Session

	query := `DELETE FROM students WHERE id = ?`

	err := session.Query(query, id).WithContext(c.Context).Exec()
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}
