package model

import "fmt"

type (
	Estudent struct {
		Id            int    `json: "id" db: "id"`
		Name          string `json: "name" db: "name"`
		Email         string `json: "email" db: "email"`
		Percentage    []float32
		Qualification float32 `json: "qualification" db: "qualification"`
	}
	Estudents []Estudent
)

func (e *Estudent) TotalQualification() float32 {
	for _, v := range e.Percentage {
		e.Qualification = e.Qualification + v
	}
	e.Qualification = e.Qualification / float32(len(e.Percentage))
	fmt.Println(e.Qualification)
	return e.Qualification
}
