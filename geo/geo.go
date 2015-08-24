package geo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	STATUS_OK               = "OK"
	STATUS_ZERO_RESULTS     = "ZERO_RESULTS"
	STATUS_OVER_QUERY_LIMIT = "OVER_QUERY_LIMIT"
	STATUS_REQUEST_DENIED   = "REQUEST_DENIED"
	STATUS_INVALID_REQUEST  = "INVALID_REQUEST"
	GOOGLE_ENDPOINT         = "https://maps.googleapis.com/maps/api/geocode/json?sensor=false&address="
)

var (
	RemoteServerError = errors.New("Unable to contact Google endpoint")
	BodyReadError     = errors.New("Unable to read the response body")
)

type Address struct {
	Lat      float64   `json:"lat"`
	Lng      float64   `json:"lng"`
	Address  string    `json:"address"`
	Response *Response `json:"response"`
}

type Response struct {
	Status  string   `json:"status"`
	Results []Result `json:"results"`
}

type Result struct {
	Types             []string           `json:"types"`
	FormattedAddress  string             `json:"formatted_address"`
	AddressComponents []AddressComponent `json:"address_components"`
	Geometry          GeometryData       `json:"geometry"`
}

type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type GeometryData struct {
	Location     LatLng `json:"location"`
	LocationType string `json:"location_type"`
	Viewport     Loc    `json:"viewport"`
	Bounds       Loc    `json:bounds"`
}

type Loc struct {
	Southwest LatLng `json:"southwest"`
	Northeast LatLng `json:"northeast"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (a *Address) String() string {
	return fmt.Sprintf("%s (lat: %3.7f, lng: %3.7f)", a.Address, a.Lat, a.Lng)
}

func Geocode(q string) (*Address, error) {
	return fetch(GOOGLE_ENDPOINT + url.QueryEscape(strings.TrimSpace(q)))
}

func ReverseGeocode(ll string) (*Address, error) {
	return fetch(GOOGLE_ENDPOINT + url.QueryEscape(strings.TrimSpace(ll)))
}

func fetch(url string) (*Address, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, RemoteServerError
	}

	defer resp.Body.Close()

	g := new(Response)
	if err := json.NewDecoder(resp.Body).Decode(g); err != nil {
		return nil, err
	}

	if g.Status != STATUS_OK {
		return nil, fmt.Errorf("Geocoder service error (%s)", g.Status)
	}

	return &Address{
		Lat:      g.Results[0].Geometry.Location.Lat,
		Lng:      g.Results[0].Geometry.Location.Lng,
		Address:  g.Results[0].FormattedAddress,
		Response: g,
	}, nil

}
