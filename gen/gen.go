//go:generate go run gen.go
//go:generate go fmt ../db.go

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	sourceData, err := os.Open("../Gaz_zcta_national.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceData.Close()

	reader := bufio.NewReader(sourceData)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	arr := [100000]string{}

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		zip, lat, long := fields[0], fields[7], fields[8]

		index, err := strconv.Atoi(zip)

		if err == nil {
			str := fmt.Sprintf("%s,%s", lat, long)
			arr[index] = str
		}
	}

	contents := []byte("package zipcode; var zipcodes = [100000]string{\n")
	for _, el := range arr {
		line := fmt.Sprintf("\"%s\",\n", el)
		contents = append(contents, line...)
	}
	contents = append(contents, "}"...)

	err = ioutil.WriteFile("../db.go", contents, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
