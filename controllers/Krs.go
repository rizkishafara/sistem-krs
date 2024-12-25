package controllers

import (
	"log"
	"net/http"
	"sistem-krs/models"

	"github.com/flosch/pongo2/v4"
	"github.com/labstack/echo/v4"
)

func KrsPage(c echo.Context) error {
	id := c.QueryParam("id")
	nama := c.QueryParam("nama")
	nim := c.QueryParam("nim")
	ipk := c.QueryParam("ipk")

	respMatkul := models.GetMatakuliah()

	if respMatkul.Status != 200 {
		log.Println("error get matakuliah", respMatkul.Message)
		return c.JSON(http.StatusInternalServerError, respMatkul)
	}
	data := pongo2.Context{
		"title":      "Krs",
		"id":         id,
		"nama":       nama,
		"nim":        nim,
		"ipk":        ipk,
		"matakuliah": respMatkul.Data,
	}
	return c.Render(http.StatusOK, "views/krs.html", data)
}
func GetKrsMhs(c echo.Context) error {
	id_mhs := c.Param("id")
	GetKrsMhs := models.GetKrsMhs(id_mhs)
	return c.JSON(http.StatusOK, GetKrsMhs)
}
func GetMatakuliah(c echo.Context) error {
	GetMatakuliah := models.GetMatakuliah()
	return c.JSON(http.StatusOK, GetMatakuliah)
}
func AddKrsMhs(c echo.Context) error {
	id_mhs := c.FormValue("id_mhs")
	id_matkul := c.FormValue("id_matkul")
	AddKrsMhs := models.AddKrsMhs(id_mhs, id_matkul)
	return c.JSON(http.StatusOK, AddKrsMhs)
}
func DeleteKrsMhs(c echo.Context) error {
	id := c.Param("id_jadwal")
	DeleteKrsMhs := models.DeletKrsMhs(id)
	return c.JSON(http.StatusOK, DeleteKrsMhs)
}
