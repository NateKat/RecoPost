package parcel

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

type Parcel struct {
	uid    string
	weight int
}

/* Global */
var Parcel_m = make(map[string]bool) // map of all parcel names

func New(scanner *bufio.Scanner) (*Parcel, error) {
	scanner.Scan()
	line := scanner.Text()

	fields := strings.Fields(line)
	if len(fields) != 2 {
		return nil, errors.New("error: parcel input should be two parameters")
	}

	if _, found := Parcel_m[fields[0]]; found {
		return nil, errors.New("error: parcel name already exist")
	}

	wt, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, err
	} else if wt < 0 {
		return nil, errors.New("error: weight should be non-negative int")
	}

	p := Parcel{fields[0], wt}
	Parcel_m[fields[0]] = true
	return &p, nil
}

func (p Parcel) Parcel_wt() int {
	return p.weight
}
