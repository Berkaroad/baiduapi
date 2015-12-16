package baiduapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherInfo struct {
	City      string  `json:"city"`
	Pinyin    string  `json:"pinyin"`
	CityCode  string  `json:"citycode"`
	Date      string  `json:"date"`
	Time      string  `json:"time"`
	PostCode  string  `json:"postCode"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Altitude  string  `json:"altitude"`
	Weather   string  `json:"weather"`
	Temp      string  `json:"temp"`
	L_Tmp     string  `json:"l_tmp"`
	H_Tmp     string  `json:"h_tmp"`
	WD        string  `json:"WD"`
	WS        string  `json:"WS"`
	Sunrise   string  `json:"sunrise"`
	Sunset    string  `json:"sunset"`
}

type WeatherInfoResult struct {
	ErrNum      int    `json:"errNum"`
	ErrMsg      string `json:"errMsg"`
	WeatherInfo `json:"retData"`
}

func (self *BaiduBiz) GetWeatherByCityName(cityName string) (WeatherInfoResult, error) {
	result := WeatherInfoResult{}
	url := "http://apis.baidu.com/apistore/weatherservice/cityname?cityname=" + cityName
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("apikey", self.ApiKey)
	if resp, err := client.Do(req); err == nil {
		if buffer, err := ioutil.ReadAll(resp.Body); err == nil {
			fmt.Println("baiduapi.BaiduBiz.GetWeatherByCityName=" + string(buffer))
			json.Unmarshal(buffer, &result)
			if result.ErrNum != 0 {
				return result, errors.New(result.ErrMsg)
			} else {
				return result, nil
			}
		}
	}
	return result, errors.New(BDUAPI_AccessServer_Error)
}
