package zipcode

import (
	"errors"
	"fmt"
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
	coord = zipcodes[zip]
	if coord.isEmpty() {
		err = errors.New("nope")
	}
	return
}
