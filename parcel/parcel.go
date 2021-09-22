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

func New(scanner *bufio.Scanner) (*Parcel, error) {
	scanner.Scan()
	line := scanner.Text()

	fields := strings.Fields(line)
	if len(fields) != 2 {
		return nil, errors.New("error: parcel input should be two parameters")
	}
	wt, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, err
	} else if wt < 0 {
		return nil, errors.New("error: weight should be non-negative int")
	}

	p := Parcel{fields[0], wt}
	return &p, nil
}
