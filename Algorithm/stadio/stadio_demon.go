package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {

	// read line
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	str := string(bytes)
	fmt.Println(str)

	// read one
	var firstName, lastName string
	fmt.Scanln(&firstName, &lastName)
	fmt.Println(firstName + lastName)
}


