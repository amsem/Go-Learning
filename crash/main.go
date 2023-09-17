
package main

import ("fmt"
	"github.com/amsem/crash/data")


var Age int = 22
func init() {
	print("I m first \n")

}

func Birthday(a *int){
	//panic("not my birthday")
	*a++
}

func main() {
	fmt.Println(entry)
	printData()
	data.Chart()
	fmt.Println(data.Y)
	Birthday(&Age)
	fmt.Printf("The memory adresse of the variable AGE is %v and the value is %v\n",&Age, Age)
}
