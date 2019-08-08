package controllers

import (
	"net/http"

	"github.com/iswanulumam/go-restful-training/library"
	"github.com/labstack/echo"
)

// get weather data
func GetWeatherController(c echo.Context) error {
	// get data from param
	city := c.Param("city")

	// get data from lib and conver kelvin to celsius
	kelvin, err := library.Weather(city)
	celsius := kelvin.(float64) - 273.15

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"City":                   city,
		"Temperature in Celsius": celsius,
	})
}
