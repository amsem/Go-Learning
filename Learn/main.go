package main

import (
	"fmt"

	"github.com/amsem/Learn/data"
)

func main() {
	var max data.Instructor
	max.Firstname = "amsem"
	max.Id = 12
	max.Lastname = "amine"
	max.Score = 100
	said := data.NewInstructor("said", "selmani", 90)
	fmt.Print(said.Print())
	fmt.Print(max.Print())
	goCourse := data.Course{Id: 2, Name: "Go Course", Instructor: said}
	fmt.Print(goCourse.String())
	swiftW := data.NewW("Swift fund", said)
	fmt.Printf(swiftW.String())
	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftW
	for _, c := range courses {
		fmt.Println(c)
	}

}
