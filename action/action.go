package action

import (
	"RecoPost/city"
	"bufio"
	"errors"
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

func (action *Action) op_one() error {
	return nil
}

func (action *Action) op_two() error {
	return nil
}

func (action *Action) op_three() error {
	//fmt.Println("Town with the most number of packages is ", city.Max_parcels_name())
	return nil
}

func Execute_actions(action_list []Action, cities_m map[string]city.City) error {
	for _, a := range action_list {
		switch a.op {
		case 1:
			a.op_one()
		case 2:
			a.op_two()
		case 3:
			a.op_three()
		default:
			return errors.New("error: action opcode doesn't exists")
		}
	}

	return nil
}
