package router

import (
	"sistem-krs/controllers"

	"github.com/labstack/echo/v4"
)

func Mahasiswa(e *echo.Echo) {
	e.GET("/", controllers.HomePage)
	e.GET("/mahasiswa", controllers.GetMhs)
	e.POST("/mahasiswa", controllers.CreateMhs)
	e.DELETE("/mahasiswa/:id", controllers.DeleteMhs)
}
func Krs(e *echo.Echo) {
	e.GET("/krs", controllers.KrsPage)
	e.GET("/krs/:id", controllers.GetKrsMhs)
	e.GET("/matakuliah", controllers.GetMatakuliah)
	e.POST("/krs", controllers.AddKrsMhs)
	e.DELETE("/krs/:id_jadwal", controllers.DeleteKrsMhs)

	e.GET("/cetak-krs", controllers.CetakKrsPage)
}
