package city

import (
	"RecoPost/office"
	"bufio"
	"errors"
	"strconv"
)

type City struct {
	cityName   string
	totParcels int
	offices    []office.Office
}

func parse_city_params(scanner *bufio.Scanner) (string, int, error) {
	scanner.Scan()
	cityName := scanner.Text()
	if len(cityName) != 1 {
		return "", 0, errors.New("error: expected city name")
	}

	scanner.Scan()
	text := scanner.Text()
	if len(text) != 1 {
		return "", 0, errors.New("error: expected number of offices")
	}

	num_offices, err := strconv.Atoi(text)
	if err != nil {
		return "", 0, err
	} else if num_offices < 0 {
		return "", 0, errors.New("error: number of offices should be non-negative int")
	}

	return cityName, num_offices, nil
}

func create_offices(scanner *bufio.Scanner, num_offices int) ([]office.Office, error) {
	var offices_list []office.Office

	for i := 0; i < num_offices; i++ {
		o, err := office.New(scanner, i)
		if err != nil {
			return nil, err
		}
		offices_list = append(offices_list, *o)
	}

	return offices_list, nil
}

func New(scanner *bufio.Scanner) (*City, error) {
	var tot = 0

	cityName, num_offices, err := parse_city_params(scanner)
	if err != nil {
		return nil, err
	}

	offices_list, err := create_offices(scanner, num_offices)
	if err != nil {
		return nil, err
	}

	for _, o := range offices_list {
		tot += o.Tot_parcels()
	}

	c := City{cityName, tot, offices_list}
	return &c, nil
}
