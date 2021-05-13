package student

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zopsmart/gofr/pkg/datastore"
	"github.com/zopsmart/gofr/pkg/gofr"
	"github.com/zopsmart/sample-crud/models"
)

func initializeTest(t *testing.T) *gofr.Gofr {
	os.Setenv("GOFR_ENV", "test")

	k := gofr.New()

	//initializing the seeder
	seeder := datastore.NewSeeder(&k.DataStore, "../../db")
	seeder.RefreshCassandra(t, "students")

	return k
}

func TestStore_Find(t *testing.T) {
	testcases := []struct {
		description string
		input       string
		output      *models.Student
		err         error
	}{
		{description: "Success Case: Find Student with ID 1", input: "1", output: &models.Student{ID: "1", Name: "Student1", Email: "student1@gmail.com", GPA: 8}},
	}

	store := New()
	k := initializeTest(t)

	for i, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, k)
		output, err := store.Find(ctx, tc.input)

		assert.Equal(t, tc.err, err, "[TEST %v], failed. %s", i+1, tc.description)

		assert.Equal(t, tc.output, output, "[TEST %v], failed. %s", i+1, tc.description)
	}
}

func TestStore_Create(t *testing.T) {
	testcases := []struct {
		description string
		input       *models.Student
	}{
		{description: "Successful Insertion", input: &models.Student{ID: "6", Name: "Student6", Email: "student6@gmail.com", GPA: 8}},
	}

	store := New()
	k := initializeTest(t)

	for i, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, k)
		err := store.Create(ctx, tc.input)

		assert.Equal(t, nil, err, "[TEST %v], failed. %s", i+1, tc.description)
	}
}

func TestStore_Update(t *testing.T) {
	testcases := []struct {
		description string
		id          string
		input       *models.Student
		err         error
	}{
		{description: "Success Case: Update Student with ID 1", id: "1", input: &models.Student{Name: "Student1", Email: "newemail@gmail.com", GPA: 8}},
	}

	store := New()
	k := initializeTest(t)

	for i, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, k)
		err := store.Update(ctx, tc.id, tc.input)

		assert.Equal(t, tc.err, err, "[TEST %v], failed. %s", i+1, tc.description)
	}
}

func TestStore_Delete(t *testing.T) {
	testcases := []struct {
		description string
		input       string
		err         error
	}{
		{description: "Success Case: Delete Student with ID 2", input: "2"},
	}

	store := New()
	k := initializeTest(t)

	for i, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, k)
		err := store.Delete(ctx, tc.input)

		assert.Equal(t, tc.err, err, "[TEST %v], failed. %s", i+1, tc.description)
	}
}
