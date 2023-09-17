package data

import "time"

type Workshop struct {
	Course 
	Date time.Time
}

func (c Course) SignUp() bool{
	return true
}
func NewW(n string, i Instructor) Workshop{
	w := Workshop {}
	w.Name = n
	w.Instructor = i
	return w
}
