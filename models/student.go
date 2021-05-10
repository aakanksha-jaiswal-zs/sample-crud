package models

type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	GPA   float64 `json:"gpa"`
}
