package models

import (
	"log"
	"sistem-krs/db"
	"sistem-krs/utils"
)

func GetMhs() utils.Respon {
	dbEngine := db.ConnectDB()
	defer dbEngine.Close() 

	var Respon utils.Respon

	query := `
		SELECT 
			m.id, m.namamhs, m.nim, m.ipk, m.sks, 
			JSON_ARRAYAGG(j.matakuliah) AS matakuliah 
		FROM 
			inputmhs m
		LEFT JOIN 
			jwl_mhs j ON m.id = j.mhs_id
		GROUP BY 
			m.id
	`
	datauser, err := dbEngine.QueryString(query)
	if err != nil {
		log.Println("error get mhs", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	Respon.Status = 200
	Respon.Data = datauser
	Respon.Message = "success"
	return Respon
}

func CreateMhs(namamhs, nim, ipk, sks, matakuliah string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	datauser, err := dbEngine.QueryString("SELECT nim FROM inputmhs WHERE nim = ? ", nim)
	if err != nil {
		log.Println("error get mhs by nim", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	defer dbEngine.Close()
	if datauser != nil {
		Respon.Status = 400
		Respon.Message = "Nim harus unik"
		return Respon
	}
	_, err = dbEngine.QueryString("INSERT INTO inputmhs (namamhs,nim,ipk,sks) VALUES (?,?,?,?)", namamhs, nim, ipk, sks)

	if err != nil {
		log.Println("error insert mahasiswa", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	defer dbEngine.Close()

	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
func DeleteMhs(id string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.QueryString("DELETE FROM inputmhs WHERE id = ?", id)
	if err != nil {
		log.Println("error delete mahasiswa", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	defer dbEngine.Close()

	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
