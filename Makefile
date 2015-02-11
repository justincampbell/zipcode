COVERAGE_FILE := coverage.out

all: dependencies db.go test

test:
	go test -coverprofile=$(COVERAGE_FILE)

dependencies:
	go get golang.org/x/tools/cmd/cover

coverage: test
	go tool cover -html=$(COVERAGE_FILE)

db.go: Gaz_zcta_national.txt
	go generate gen/*.go

Gaz_zcta_national.txt: Gaz_zcta_national.zip
	unzip -o Gaz_zcta_national.zip
	touch Gaz_zcta_national.txt

Gaz_zcta_national.zip:
	wget http://www.census.gov/geo/maps-data/data/docs/gazetteer/Gaz_zcta_national.zip

.PHONY: all coverage dependencies test
