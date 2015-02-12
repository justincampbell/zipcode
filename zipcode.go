package zipcode

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Coordinate struct {
	lat  string
	long string
}

func (coord *Coordinate) isEmpty() bool {
	return coord.lat == "" && coord.long == ""
}

func (coord *Coordinate) String() string {
	if coord.isEmpty() {
		return ""
	}
	return fmt.Sprintf("%s,%s", coord.lat, coord.long)
}

func Lookup(zip string) (coord Coordinate, err error) {
	index, _ := strconv.Atoi(zip)
	el := zipcodes[index]
	if el == "" {
		err = errors.New(fmt.Sprintf("Could not find ZIP Code %s", zip))
		return
	}
	parts := strings.Split(el, ",")
	coord = Coordinate{lat: parts[0], long: parts[1]}
	return
}
