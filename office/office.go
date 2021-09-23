package office

import (
	"RecoPost/parcel"
	"bufio"
	"errors"
	"strconv"
	"strings"
)

type Office struct {
	officeName  int
	maxParcelWt int
	minParcelWt int
	parcels     []parcel.Parcel
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func parse_office_params(line string) ([]int, error) {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		return nil, errors.New("error: office input should be three parameters")
	}

	fields_i, err := sliceAtoi(fields)
	if err != nil {
		return nil, err
	}
	for _, number := range fields_i {
		if number < 0 {
			return nil, errors.New("error: office parameters should be non-negative int")
		}
	}
	if fields_i[2] < fields_i[1] {
		return nil, errors.New("error: parcel weight, minimum cannot exceed maximum")
	}

	return fields_i, nil
}

func create_parcels(scanner *bufio.Scanner, o_params []int) ([]parcel.Parcel, error) {
	var parcels_list []parcel.Parcel

	for i := 0; i < o_params[0]; i++ {
		p, err := parcel.New(scanner)
		if err != nil {
			return nil, err
		}
		if o_params[1] <= p.Parcel_wt() && p.Parcel_wt() <= o_params[2] {
			parcels_list = append(parcels_list, *p)
		} else {
			return nil, errors.New("error: parcel weight is not within office limit")
		}
	}

	return parcels_list, nil
}

func New(scanner *bufio.Scanner, office_number int) (*Office, error) {
	scanner.Scan()
	line := scanner.Text()
	o_params, err := parse_office_params(line)
	if err != nil {
		return nil, err
	}

	parcels_list, err := create_parcels(scanner, o_params)
	if err != nil {
		return nil, err
	}

	o := Office{office_number, o_params[1], o_params[2], parcels_list}
	return &o, nil
}

func (o Office) Tot_parcels() int {
	return len(o.parcels)
}

func (send_office *Office) Send_to_office(recv_office Office) (int, int) {
	var rejected_list []parcel.Parcel
	var tot_send, tot_recv int

	for _, p := range send_office.parcels {
		if recv_office.minParcelWt <= p.Parcel_wt() && p.Parcel_wt() <= recv_office.maxParcelWt {
			recv_office.parcels = append(recv_office.parcels, p) // Enqueue
			tot_recv += 1
		} else {
			rejected_list = append(rejected_list, p)
			tot_send += 1
		}
	}
	//send_office.parcels = nil
	send_office.parcels = rejected_list
	return tot_send, tot_recv
}
