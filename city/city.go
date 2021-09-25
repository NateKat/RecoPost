package city

import (
	"RecoPost/office"
	"bufio"
	"errors"
	"strconv"
)

type City struct {
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

func New(scanner *bufio.Scanner) (*City, string, error) {
	var tot = 0

	cityName, num_offices, err := parse_city_params(scanner)
	if err != nil {
		return nil, "", err
	}

	offices_list, err := create_offices(scanner, num_offices)
	if err != nil {
		return nil, "", err
	}

	for _, o := range offices_list {
		tot += o.Tot_parcels()
	}

	c := City{tot, offices_list}
	return &c, cityName, nil
}

func Move_parcels(c1, c2 City, o1, o2 int) {
	tot_send, tot_recv := c1.offices[o1].Send_to_office(c2.offices[o2])
	c1.totParcels -= tot_send
	c2.totParcels += tot_recv
}

func Max_parcels_name(cities_m map[string]City) string {
	max_t := -1
	max_c := ""

	for name, c := range cities_m {
		if max_t < c.totParcels {
			max_t = c.totParcels
			max_c = name
		}
	}

	return max_c
}

func Create_cities(scanner *bufio.Scanner, num_cities int) (map[string]City, error) {
	var cities_m = make(map[string]City)

	for i := 0; i < num_cities; i++ {
		c, name, err := New(scanner)
		if err != nil {
			return nil, err
		}
		cities_m[name] = *c
	}

	return cities_m, nil
}
