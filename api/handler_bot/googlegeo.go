package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/mingrammer/meetup-api/config"
)

type GeoCode struct {
	Lat float64
	Lng float64
}

type GeoInfo struct {
	Geometry struct {
		Location GeoCode
	}
}

type GeoAPIResponse struct {
	Results []GeoInfo
}

func encodeURL(urlStr string) string {
	return url.QueryEscape(urlStr)
}

func getGeoCodeRequestURL(location string) string {
	url := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s",
		encodeURL(location),
		config.BotAPIConfig.GoogleAPIKey,
	)
	return url
}

func getGeoCode(url string) GeoCode {
	resp, err := http.Get(url)
	if err != nil {
		return GeoCode{}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GeoCode{}
	}
	geoAPIResp := GeoAPIResponse{}
	json.Unmarshal(body, &geoAPIResp)
	return geoAPIResp.Results[0].Geometry.Location
}

func getMapImageURL(location string) string {
	reqURL := getGeoCodeRequestURL(location)
	geoCode := getGeoCode(reqURL)
	url := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/staticmap?center=%f,%f&zoom=17&size=500x500&key=%s&sense=false&markers=%f,%f",
		geoCode.Lat,
		geoCode.Lng,
		config.BotAPIConfig.GoogleMapAPIKey,
		geoCode.Lat,
		geoCode.Lng,
	)
	return url
}
