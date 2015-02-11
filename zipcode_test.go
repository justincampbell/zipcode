package zipcode

import (
	"errors"
	"testing"

	"github.com/bmizerany/assert"
)

func Test_Lookup(t *testing.T) {
	zeroes, err := Lookup("00000")
	assert.Equal(t, Coordinates{lat: "", long: ""}, zeroes)
	assert.Equal(t, errors.New("nope"), err)

	cherry_hill, err := Lookup("08003")
	assert.Equal(t, Coordinates{lat: "39.882703", long: "-74.972036"}, cherry_hill)
	assert.Equal(t, nil, err)
}
