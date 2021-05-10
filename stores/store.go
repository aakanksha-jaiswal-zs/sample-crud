package stores

import (
	"github.com/zopsmart/gofr/pkg/errors"
	"github.com/zopsmart/gofr/pkg/gofr"
	"github.com/zopsmart/sample-crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type student struct {
}

func New() *student {
	return &student{}
}

func (s *student) Find(c *gofr.Context, id string) (*models.Student, error) {
	collection := c.MongoDB.Collection("students")

	res := collection.FindOne(c.Context, bson.D{primitive.E{Key: "id", Value: id}})
	if err := res.Err(); err != nil {
		return nil, err
	}

	var student models.Student

	err := res.Decode(&student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *student) Create(c *gofr.Context, student *models.Student) error {
	collection := c.MongoDB.Collection("students")

	_, err := collection.InsertOne(c.Context, student)

	return err
}

func (s *student) Update(c *gofr.Context, id string, student *models.Student) error {
	collection := c.MongoDB.Collection("students")

	res := collection.FindOneAndUpdate(c.Context, bson.D{primitive.E{Key: "id", Value: id}}, getUpdate(student))
	if res == nil {
		return errors.DB{} //todo remove
	}

	return res.Err()
}

func getUpdate(student *models.Student) bson.M {
	update := bson.M{}
	if student.Name != "" {
		update["name"] = student.Name
	}

	if student.GPA != 0 {
		update["gpa"] = student.GPA
	}

	if student.Email != "" {
		update["email"] = student.Email
	}

	return bson.M{"$set": update}
}

func (s *student) Delete(c *gofr.Context, id string) error {
	collection := c.MongoDB.Collection("students")

	_, err := collection.DeleteOne(c.Context, bson.D{primitive.E{Key: "id", Value: id}})
	return err
}
