package main 

import "fmt"

func main(){
	var opt string
	var num1, num2 float64
	for {
	fmt.Println("Welcome to AMSEM CALCULATTOR !!!")
       	fmt.Println("================================")
	fmt.Println("Which operation you need to do ? ===> (add, sub, mult, div)")
	fmt.Scanf("%s", &opt)
	fmt.Print("Number 1 = ")
	fmt.Scanf("%f", &num1)
	fmt.Print("Number 2 = ")
	fmt.Scanf("%f", &num2)

	switch opt {
		case "add":
			fmt.Println(num1 + num2)
		case "sub":
                        fmt.Println(num1 - num2)
		case "mult":
                        fmt.Println(num1 * num2)
		case "div":
                        fmt.Println(num1 / num2)
	}
	}
}
