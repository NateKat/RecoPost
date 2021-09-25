package main

import (
	"RecoPost/action"
	"RecoPost/city"
	"RecoPost/parcel"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func parse_single_number(scanner *bufio.Scanner) (int, error) {
	scanner.Scan()
	text := scanner.Text()
	if len(text) != 1 {
		return 0, errors.New("error: expected single number")
	}

	number, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	} else if number < 0 {
		return 0, errors.New("error: number should be non-negative int")
	}

	return number, nil
}

func handle_cities_input(scanner *bufio.Scanner) (map[string]*city.City, error) {
	var parcel_m = parcel.Parcel_m // map[string]bool

	num_cities, err := parse_single_number(scanner)
	if err != nil {
		return nil, err
	}

	cities_m, err := city.Create_cities(scanner, num_cities)
	if err != nil {
		return nil, err
	}

	for k := range parcel_m {
		delete(parcel_m, k) // free the map when creation is done
	}

	return cities_m, nil
}

func actions_input_and_exec(scanner *bufio.Scanner, cities_m map[string]*city.City) error {
	num_actions, err := parse_single_number(scanner)
	if err != nil {
		return err
	}

	action_list, err := action.Create_actions(scanner, num_actions)
	if err != nil {
		return err
	}

	return action.Execute_actions(action_list, cities_m)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cities_m, err := handle_cities_input(scanner)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	err = actions_input_and_exec(scanner, cities_m)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if scanner.Err() != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
