package main

import (
	"RecoPost/city"
	"RecoPost/parcel"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func parse_num_cities(scanner *bufio.Scanner) (int, error) {
	scanner.Scan()
	text := scanner.Text()
	if len(text) != 1 {
		return 0, errors.New("error: expected number of cities")
	}

	num_cities, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	} else if num_cities < 0 {
		return 0, errors.New("error: number of cities should be non-negative int")
	}

	return num_cities, nil
}

func create_cities(scanner *bufio.Scanner, num_cities int) ([]city.City, error) {
	var cities_list []city.City

	for i := 0; i < num_cities; i++ {
		c, err := city.New(scanner)
		if err != nil {
			return nil, err
		}
		cities_list = append(cities_list, *c)
	}

	return cities_list, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var parcel_m = parcel.Parcel_m // map[string]bool

	num_cities, err := parse_num_cities(scanner)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	cities_list, err := create_cities(scanner, num_cities)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	for k := range parcel_m {
		delete(parcel_m, k) // free the map when creation is done
	}

	if scanner.Err() != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println("DONE", cities_list, "is the city list") // Debug
}
