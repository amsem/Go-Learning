package main

import ("fmt"
		"bufio"
		"os"
		//"strconv"
	)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	fmt.Printf("you typed %q\n", input)
}