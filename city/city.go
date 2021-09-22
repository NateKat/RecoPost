package city

import (
	"fmt"
	"bufio"
	"strconv"
	"RecoPost/office"
)

type City struct {
	cityName   string
	totParcels int
	offices    []office.Office
}

func New(cityName string, scanner *bufio.Scanner) City {
	var	offices_list    []office.Office

	fmt.Print("Offices in city: ")
	scanner.Scan() // Scans a line from Stdin(Console)

	text := scanner.Text()
	if len(text) == 0 {
		fmt.Println("Error: empty input")
	}
	fmt.Println(text, "offices") // Debug
	num_offices, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < num_offices; i++ {
		o := office.New(scanner)
		offices_list = append(offices_list, o)
	}
	c := City{cityName, 0, offices_list}
	return c
}

/*
func (c City) update_city(scanner *bufio.Scanner) {
	fmt.Print("Offices in city: ")
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
}
*/

// totParcels int, offices []office.Office
