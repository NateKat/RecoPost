package office

import (
	"fmt"
	"bufio"
	"strconv"
	"RecoPost/parcel"
)

type Office struct {
	officeName  string
	maxParcelWt int
	minParcelWt int
	parcels     []parcel.Parcel
}

func New(scanner *bufio.Scanner) Office {
	// officeName string, maxParcelWt int, minParcelWt int, parcellist []parcel.Parcel
	var parcels_list []parcel.Parcel
	fmt.Print("parcels and wt_range in office: ")
	scanner.Scan() // Scans a line from Stdin(Console)

	text := scanner.Text()
	if len(text) == 0 {
		fmt.Println("Error: empty input")
	}
	fmt.Println(text, "parcels and wt:") // Debug
	num_parcels, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < num_parcels; i++ {
		fmt.Print("each parcel name and wt: ")
		scanner.Scan() // Scans a line from Stdin(Console)

		text := scanner.Text()
		if len(text) == 0 {
			fmt.Println("Error: empty input")
		}
		fmt.Println(text, "specific parcel:") // Debug
		p := parcel.New(text, 0)
		parcels_list = append(parcels_list, p)

	}

	o := Office{"0", 0, 0, parcels_list}
	return o
}
