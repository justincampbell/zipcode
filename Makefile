all: dependencies db.go test

test:
	go test

dependencies:
	go get -t

db.go: Gaz_zcta_national.txt
	go generate gen/*.go

Gaz_zcta_national.txt: Gaz_zcta_national.zip
	unzip -o Gaz_zcta_national.zip
	touch Gaz_zcta_national.txt

Gaz_zcta_national.zip:
	wget http://www.census.gov/geo/maps-data/data/docs/gazetteer/Gaz_zcta_national.zip

.PHONY: all dependencies test
