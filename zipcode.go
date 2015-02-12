package zipcode

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	Lat  string
	Long string
}

func (coord *Coordinate) isEmpty() bool {
	return coord.Lat == "" && coord.Long == ""
}

func (coord *Coordinate) String() string {
	if coord.isEmpty() {
		return ""
	}
	return fmt.Sprintf("%s,%s", coord.Lat, coord.Long)
}

func Lookup(zip string) (coord Coordinate, err error) {
	index, _ := strconv.Atoi(zip)
	el := zipcodes[index]
	if el == "" {
		err = errors.New(fmt.Sprintf("Could not find ZIP Code %s", zip))
		return
	}
	parts := strings.Split(el, ",")
	coord = Coordinate{Lat: parts[0], Long: parts[1]}
	return
}
