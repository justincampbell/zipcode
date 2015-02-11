package zipcode

import "testing"

func Test_Lookup_zeroes(t *testing.T) {
	coord, err := Lookup("00000")

	if !coord.isEmpty() {
		t.Fatalf("coord: %#v", coord)
	}

	if err == nil {
		t.Fatal("should be err")
	}
}

func Test_Lookup_cherry_hill(t *testing.T) {
	coord, err := Lookup("08003")

	if coord != (Coordinate{lat: "39.882703", long: "-74.972036"}) {
		t.Fatalf("coord: %#v", coord)
	}

	if err != nil {
		t.Fatalf("err: %#v", err)
	}
}
