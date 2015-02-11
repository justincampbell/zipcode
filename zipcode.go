package zipcode

import "errors"

type Coordinate string

type Coordinates struct {
	lat  Coordinate
	long Coordinate
}

func (coord *Coordinates) isEmpty() bool {
	return coord.lat == "" && coord.long == ""
}

func Lookup(zip string) (coord Coordinates, err error) {
	coord = zipcodes[zip]
	if coord.isEmpty() {
		err = errors.New("nope")
	}
	return
}
