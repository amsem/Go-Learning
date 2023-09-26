package data

import "fmt"

func NewInstructor(f string, l string, s int) Instructor {
	return Instructor{Firstname: f, Lastname: l, Score: s}
}

func (i Instructor) Print() string {
	return fmt.Sprintf("%v, %v (%d)\n", i.Firstname, i.Lastname, i.Score)
}

func (c Course) String() string {
	return fmt.Sprintf("===>%v\n====>%v\n", c.Name, c.Instructor.Lastname)
}
