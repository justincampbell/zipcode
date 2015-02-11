package zipcode

import "errors"

type Coordinate struct {
	lat  string
	long string
}

func (coord *Coordinate) isEmpty() bool {
	return coord.lat == "" && coord.long == ""
}

func Lookup(zip string) (coord Coordinate, err error) {
	coord = zipcodes[zip]
	if coord.isEmpty() {
		err = errors.New("nope")
	}
	return
}
