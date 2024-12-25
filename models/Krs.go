package models

import (
	"log"
	"sistem-krs/db"
	"sistem-krs/utils"
)

func GetKrsMhs(id_mhs string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	datakrs, err := dbEngine.QueryString("SELECT * FROM jwl_mhs WHERE mhs_id = ? ", id_mhs)
	if err != nil {
		log.Println("error get krs", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	defer dbEngine.Close()

	Respon.Status = 200
	Respon.Data = datakrs
	Respon.Message = "success"
	return Respon
}
func GetMatakuliah() utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	datakrs, err := dbEngine.QueryString("SELECT * FROM jwl_matakuliah")
	if err != nil {
		log.Println("error get krs", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	defer dbEngine.Close()

	Respon.Status = 200
	Respon.Data = datakrs
	Respon.Message = "success"
	return Respon
}
func AddKrsMhs(id_mhs, id_matkul string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	datamatakul, err := dbEngine.QueryString("SELECT matakuliah, sks, kelp, ruangan FROM jwl_matakuliah WHERE id = ? ", id_matkul)
	if err != nil {
		log.Println("error get krs", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	defer dbEngine.Close()

	matakuliah := datamatakul[0]["matakuliah"]
	sks := datamatakul[0]["sks"]
	kelp := datamatakul[0]["kelp"]
	ruangan := datamatakul[0]["ruangan"]

	_, err = dbEngine.QueryString("INSERT INTO jwl_mhs (mhs_id, matakuliah, sks, kelp, ruangan) VALUES (?,?,?,?,?)", id_mhs, matakuliah, sks, kelp, ruangan)

	if err != nil {
		log.Println("error insert krs", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	defer dbEngine.Close()

	//append matakuliah in inputmhs table
	_, err = dbEngine.QueryString("UPDATE inputmhs SET matakuliah = CONCAT(matakuliah, ?) WHERE id = ?", id_matkul+", ", id_mhs)
	if err != nil {
		log.Println("error updating inputmhs", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}

	Respon.Status = 200
	Respon.Message = "Berhasil menambah matkul"

	return Respon
}
func DeletKrsMhs(id_jadwal string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	_, err := dbEngine.QueryString("DELETE FROM jwl_mhs WHERE id = ?", id_jadwal)
	if err != nil {
		log.Println("error delete krs", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	defer dbEngine.Close()

	Respon.Status = 200
	Respon.Message = "Berhasil menghapus krs"
	return Respon
}
