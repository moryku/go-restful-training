package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// A Response struct to map the Entire Response
type Response struct {
	Main Main `json:"main"`
}

type Main struct {
	Temp float64 `json:"temp"`
}

func main() {
	var api_key string = "b63ab79141550fee57c0c6e64e5132a2"
	var url_api string = "https://api.openweathermap.org/data/2.5/weather?" + "appid=" + api_key + "&q=" + "Jakarta"

	resp, err := http.Get(url_api)
	if err != nil {
		fmt.Println(err.Error())
	}
	respRead, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()

	var weather Response
	json.Unmarshal(respRead, &weather)
	fmt.Println(weather.Main.Temp)
}
