package zipcode

import (
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	Lat  float64
	Long float64
}

func (coord *Coordinate) isEmpty() bool {
	return coord.Lat == 0 && coord.Long == 0
}

func (coord *Coordinate) String() string {
	if coord.isEmpty() {
		return ""
	}
	return fmt.Sprintf("%f,%f", coord.Lat, coord.Long)
}

func Lookup(zip string) (coord Coordinate, err error) {
	index, _ := strconv.Atoi(zip)
	el := zipcodes[index]
	if el == "" {
		err = fmt.Errorf("Could not find ZIP Code %s", zip)
		return
	}
	parts := strings.Split(el, ",")
	lat, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return
	}
	long, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return
	}
	coord = Coordinate{Lat: lat, Long: long}
	return
}
