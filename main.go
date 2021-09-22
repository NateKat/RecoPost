package main

import (
	"RecoPost/city"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var cities_list []city.City
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
		c, err := city.New(scanner)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		cities_list = append(cities_list, *c)
	}
	if scanner.Err() != nil {
		fmt.Print("TODO: Handle error")
	}
	fmt.Println("DONE", cities_list, "is the city list") // Debug
}
