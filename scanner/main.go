package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Print("Please enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()
	fmt.Println("Hello", name)

}
