package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Num of cities: ")
	scanner.Scan() // Scans a line from Stdin(Console)

	text := scanner.Text()
	if len(text) == 0 {
		fmt.Println("Error: empty input")
	}
	fmt.Println(text, "cities") // Debug

	num_cities, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < num_cities; i++ {
		// TODO: new city
	}
	if scanner.Err() != nil {
		// Handle error.
	}

}
