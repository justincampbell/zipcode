//go:generate go run gen.go
//go:generate go fmt ../db.go

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	contents := []byte("package zipcode; var zipcodes = map[string]Coordinate{")

	sourceData, err := os.Open("../Gaz_zcta_national.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceData.Close()

	reader := bufio.NewReader(sourceData)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		zip, lat, long := fields[0], fields[7], fields[8]

		if zip == "GEOID" {
			continue
		}

		line := fmt.Sprintf("\n\"%s\": Coordinate{lat: \"%s\", long: \"%s\"},", zip, lat, long)

		contents = append(contents, line...)
	}

	contents = append(contents, "}"...)

	err = ioutil.WriteFile("../db.go", contents, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
