package action

import (
	"RecoPost/city"
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Action struct {
	op   int
	args string
}

func verify_int(str string) error {
	num, err := strconv.Atoi(str)
	if err != nil {
		return err
	} else if num < 0 {
		return errors.New("error: number should be non-negative int")
	}

	return nil
}

func parse_op_one(fields []string) (int, string, error) {
	if len(fields) != 2 {
		return 0, "", errors.New("error: op 1 should have 1 param")
	}

	return 1, fields[1], nil
}

func parse_op_two(fields []string) (int, string, error) {
	if len(fields) != 5 {
		return 0, "", errors.New("error: op 2 should have 4 params")
	}
	err := verify_int(fields[2])
	if err != nil {
		return 0, "", err
	}
	err = verify_int(fields[4])
	if err != nil {
		return 0, "", err
	}

	return 2, strings.Join(fields, " ")[1:], nil
}

func parse_op_three(fields []string) (int, string, error) {
	if len(fields) != 1 {
		return 0, "", errors.New("error: op 3 should have 0 params")
	}
	return 3, "", nil
}

func parse_action_params(scanner *bufio.Scanner) (int, string, error) {
	scanner.Scan()
	line := scanner.Text()

	fields := strings.Fields(line)
	if len(fields) == 0 {
		return 0, "", errors.New("error: action input should not be empty")
	}

	switch fields[0] {
	case "1":
		return parse_op_one(fields)
	case "2":
		return parse_op_two(fields)
	case "3":
		return parse_op_three(fields)
	default:
		return 0, "", errors.New("error: opcode doesn't exists")
	}
}

func New(scanner *bufio.Scanner) (*Action, error) {
	op, args, err := parse_action_params(scanner)
	if err != nil {
		return nil, err
	}

	a := Action{op, args}
	return &a, nil
}

func Create_actions(scanner *bufio.Scanner, num_actions int) ([]Action, error) {
	var action_list []Action

	for i := 0; i < num_actions; i++ {
		c, err := New(scanner)
		if err != nil {
			return nil, err
		}
		action_list = append(action_list, *c)
	}

	return action_list, nil
}

/* Print offices and parcel names for specific city */
func (action *Action) op_one(cities_m map[string]*city.City) error {
	if _, ok := cities_m[action.args]; ok {
		fmt.Print(action.args, ":\n")
		return cities_m[action.args].Print_city()
	} else {
		return errors.New("error: city name doesn't exist")
	}
}

/* Send parcels from (city, office) to different (city, office)*/
func (action *Action) op_two(cities_m map[string]*city.City) error {
	fields := strings.Fields(action.args)

	o1, err := strconv.Atoi(fields[1])
	if err != nil {
		panic("Trying to execute op with illegal value")
	}
	o2, err := strconv.Atoi(fields[3])
	if err != nil {
		panic("Trying to execute op with illegal value")
	}

	if _, ok := cities_m[fields[0]]; !ok {
		return errors.New("error: city name doesn't exist")
	}
	if _, ok := cities_m[fields[2]]; !ok {
		return errors.New("error: city name doesn't exist")
	}

	return city.Move_parcels(cities_m, fields[0], fields[2], o1, o2)
}

/* Print the name of the city with the maximum number of parcels */
func (action *Action) op_three(cities_m map[string]*city.City) error {
	fmt.Println("Town with the most number of packages is", city.Max_parcels_name(cities_m))
	return nil
}

func Execute_actions(action_list []Action, cities_m map[string]*city.City) error {
	for _, a := range action_list {
		var err error

		switch a.op {
		case 1:
			err = a.op_one(cities_m)
		case 2:
			err = a.op_two(cities_m)
		case 3:
			err = a.op_three(cities_m)
		default:
			err = errors.New("error: action opcode doesn't exists")
		}

		if err != nil {
			return err
		}
	}

	return nil
}
