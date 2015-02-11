package zipcode

import "errors"

type Coordinate string

type Coordinates struct {
	lat  Coordinate
	long Coordinate
}

type ZipCode string

func Lookup(zip ZipCode) (coord Coordinates, err error) {
	coord = zipcodes[zip]
	if coord.lat == "" && coord.long == "" {
		err = errors.New("nope")
	}
	return
}
