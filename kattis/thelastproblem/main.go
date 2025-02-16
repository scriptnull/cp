package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var name []byte
	for scanner.Scan() {
		name = append(name, scanner.Bytes()...)
	}
	fmt.Printf("Thank you, %s, and farewell!", string(name))
}
