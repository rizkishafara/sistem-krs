package models

import (
	"encoding/json"
	"log"
	"sistem-krs/db"
	"sistem-krs/utils"
)

func GetMhs() utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon
	datauser, err := dbEngine.QueryString("SELECT id, namamhs, nim, ipk, sks FROM inputmhs")
	if err != nil {
		log.Println("error get mhs", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	defer dbEngine.Close()

	for i := 0; i < len(datauser); i++ {
		id := datauser[i]["id"]
		datakrs, err := dbEngine.QueryString("SELECT matakuliah FROM jwl_mhs WHERE mhs_id = ? ", id)
		if err != nil {
			log.Println("error get krs", err)
			Respon.Status = 500
			Respon.Message = err.Error()
			return Respon
		}
		defer dbEngine.Close()
		datakrsJSON, _ := json.Marshal(datakrs)
		datauser[i]["matakuliah"] = string(datakrsJSON)
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
