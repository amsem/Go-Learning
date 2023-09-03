package go_task

import ("fmt"
	"time"
	)


type Task struct {
	ID int
	Title string
	Description string
	DueDate time.Time
	IsComplete bool
}

func CreateTask (){}
func GetTaskByID (){}
func GetAllTasks (){}
func UpdateTask () {}
func DeleteTask () {}
func MarkAsComplete () {}
func ValidateTask () {}

