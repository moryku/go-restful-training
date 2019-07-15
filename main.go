package main
import (
    "net/http"
    "github.com/labstack/echo"
)

func controllerUserWitQueryParam(c echo.Context) error {
    // url param
    urlParam := c.Param("name_param")
    // query param
    ageUser := c.QueryParam("age")
    return c.JSON(http.StatusOK, map[string]string{
        "name": urlParam,
        "age":  ageUser,
    })
}

func main() {
    e := echo.New()
    e.GET("/user/:name_param", controllerUserWitQueryParam)
    e.Start(":8000")
}