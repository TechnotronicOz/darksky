package main

import (
	"fmt"
	"github.com/technotronicoz/darksky/darksky"
	"github.com/technotronicoz/darksky/geo"
	// "io/ioutil"
	// "log"
	// "strings"
	"strconv"
)

const KEY = "9d196fe98f3eb7ed191858c4b0d735f8"

type latlng struct {
	Lat string
	Lng string
}

func main() {
	//key := strings.TrimSpace(KEY)

	location, err := geo.Geocode("3912 NW 72nd St, Kansas City MO")
	if err != nil {
		fmt.Printf("Error %+v\n", err)
	}

	// fmt.Printf("location %+v\n", location)

	lat := strconv.FormatFloat(location.Lat, 'f', 7, 64)
	lng := strconv.FormatFloat(location.Lng, 'f', 7, 64)
	// fmt.Println("lat", lat)
	// fmt.Println("lng", lng)
	// addr := location.Address

	f, err := darksky.Get(KEY, lat, lng, "now", darksky.US)
	if err != nil {
		fmt.Printf("error %+v\n", err)
	}

	fmt.Printf("Api calls: %+v\n", f.APICalls)
	fmt.Printf("Currently %s\n", f.Currently.Summary)
	fmt.Printf("Temp %.2f F\n", f.Currently.Temperature)
	fmt.Printf("Windspeed %.2f\n", f.Currently.WindSpeed)

	// fmt.Printf("%+v\n", location.Response)

	// fmt.Printf("Location %+v\n", location)
}
