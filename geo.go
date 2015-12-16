package baiduapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type AddressComponent struct {
	City          string `json:"city"`
	Direction     string `json:"direction"`
	Distance      string `json:"distance"`
	District      string `json:"district"`
	Province      string `json:"province"`
	Street        string `json:"street"`
	Street_Number string `json:"street_number"`
}

type Location struct {
	Latitude  float64 `json:"lng"`
	Longitude float64 `json:"lat"`
}

type GeographyInfo struct {
	Location          `json:"location"`
	Formatted_Address string `json:"formatted_address"`
	Business          string `json:"business"`
	AddressComponent  `json:"addressComponent"`
	CityCode          int `json:"cityCode"`
}

type GeographyInfoResult struct {
	Status        string `json:"status"`
	GeographyInfo `json:"result"`
}

func (self *BaiduBiz) GetGeographyInfo(latitude float64, longitude float64) (GeographyInfoResult, error) {
	geographyInfoResult := GeographyInfoResult{}
	url := "http://api.map.baidu.com/geocoder?location=" + strconv.FormatFloat(latitude, 'f', -1, 64) + "," + strconv.FormatFloat(longitude, 'f', -1, 64) + "&output=json&key=" + self.ApiKey
	if resp, err := http.Get(url); err == nil {
		if buffer, err := ioutil.ReadAll(resp.Body); err == nil {
			fmt.Println("baiduapi.GetLocationInfo=" + string(buffer))
			json.Unmarshal(buffer, &geographyInfoResult)
			if geographyInfoResult.Status != "OK" {
				return geographyInfoResult, errors.New(geographyInfoResult.Status)
			} else {
				return geographyInfoResult, nil
			}
		}
	}
	return geographyInfoResult, errors.New(BDUAPI_AccessServer_Error)
}
