package teacher

import "fmt"

type Teacher struct {
	Id         int
	Name       string
	Subject    string
	BaseSalary float64
	Bonus      float64
}

func (t Teacher) GetInfo () string {
	return fmt.Sprintf("Id: %d | Ten: %s | Mon hoc: %s | Luong: %.2f", t.Id, t.Name, t.Subject, t.CaculateSalary())
}

func (t Teacher) CaculateSalary() float64 {
	return t.BaseSalary + t.Bonus
}

func (t Teacher) GetId() int {
	return t.Id
}