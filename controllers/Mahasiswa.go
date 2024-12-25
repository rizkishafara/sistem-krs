package controllers

import (
	"fmt"
	"net/http"
	"sistem-krs/models"
	"strconv"

	"github.com/flosch/pongo2/v4"
	"github.com/labstack/echo/v4"
)

func HomePage(c echo.Context) error {
	fmt.Println("Home Page")
	data := pongo2.Context{
		"title": "Home",
	}
	return c.Render(http.StatusOK, "views/mahasiswa.html", data)
}
func GetMhs(c echo.Context) error {
	GetMhs := models.GetMhs()
	return c.JSON(http.StatusOK, GetMhs)
}
func CreateMhs(c echo.Context) error {
	namamhs := c.FormValue("namamhs")
	nim := c.FormValue("nim")
	ipk := c.FormValue("ipk")

	// ipk string to float
	ipkFloat, err := strconv.ParseFloat(ipk, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid IPK format",
		})
	}

	var sks string

	if ipkFloat > 4 || ipkFloat < 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Status":  "400",
			"Message": "IPK harus diantara 0-4",
		})

	} else if ipkFloat < 3.0 {
		sks = "20"
	} else if ipkFloat > 3.0 {
		sks = "24"
	}

	ipk = fmt.Sprintf("%.2f", ipkFloat)

	matakuliah := c.FormValue("matakuliah")
	CreateMhs := models.CreateMhs(namamhs, nim, ipk, sks, matakuliah)
	return c.JSON(http.StatusOK, CreateMhs)
}
func DeleteMhs(c echo.Context) error {
	id := c.Param("id")
	DeleteMhs := models.DeleteMhs(id)
	return c.JSON(http.StatusOK, DeleteMhs)
}
