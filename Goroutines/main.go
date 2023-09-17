package main

import ("fmt"
	"time"
	)


func PrintMessage(msg string) {
	for i:=0; i<5; i++{
		fmt.Println(msg)
		time.Sleep(700 * time.Millisecond)
	}
}


func main() {
	go PrintMessage("GOLANG is fantastic!!")
	go PrintMessage("Python is SHIT!!!!!!!")
	for {}
}
