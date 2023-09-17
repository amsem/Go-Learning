package main

import ("github.com/amsem/file/utils"
	"fmt"
	"os"
)
type Kdistance float64
type location string

func (origin location) DistanceTo(dest location) Kdistance {
        //TODO calc ......
        fmt.Printf("Origin %v Destination %v\n", origin, dest)
}


func main() {
	rootPath, _ := os.Getwd()
	filepath := rootPath + "/amsem.txt"
	c, _ := utils.ReadFile(filepath)
	fmt.Println(c)
	newContent:= fmt.Sprintf("Original : %v Double Original : %v \n %v\n", c,c,c)
	utils.FileWriter(rootPath + "/output.txt", newContent)
	//check the official library
	nyc := location("21.321, 312.3412")
        nyc.DistanceTo(location("32.3241, 43.43213"))

}
