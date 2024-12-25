package models

import (
	"log"
	"sistem-krs/db"
	"sistem-krs/utils"
)

func Login(username, password string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	datauser, err := dbEngine.QueryString("SELECT id, username, password FROM users WHERE username = ? ", username)

	if err != nil {
		log.Println("error get user", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	defer dbEngine.Close()

	if datauser == nil {
		Respon.Status = 404
		Respon.Message = "user not found"
		return Respon
	}

	errc := utils.CheckPasswordHash(password, datauser[0]["password"])
	if !errc {
		Respon.Status = 404
		Respon.Message = "password not match"
		return Respon
	}

	datares := make(map[string]interface{})
	if datauser != nil {
		datares["username"] = datauser[0]["username"]
		datares["id"] = datauser[0]["id"]
	}
	Respon.Status = 200
	Respon.Data = datares
	Respon.Message = "success"
	return Respon
}
func Register(username, password string) utils.Respon {
	dbEngine := db.ConnectDB()
	var Respon utils.Respon

	datauser, err := dbEngine.QueryString("SELECT id FROM users WHERE username = ? ", username)
	if err != nil {
		log.Println("error get user", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	defer dbEngine.Close()
	if datauser != nil {
		Respon.Status = 400
		Respon.Message = "username already exist"
		return Respon
	}

	passnew, err := utils.HashPassword(password)
	if err != nil {
		log.Println("error hash password", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	_, err = dbEngine.QueryString("INSERT INTO users (username, password) VALUES (?,?)", username, passnew)

	if err != nil {
		log.Println("error insert user", err)
		Respon.Status = 500
		Respon.Message = err.Error()
		return Respon
	}
	defer dbEngine.Close()

	Respon.Status = 200
	Respon.Message = "success"
	return Respon
}
