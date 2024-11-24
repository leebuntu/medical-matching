package main

import (
	"log"
	"medical-matching/controller/hospital"
	"medical-matching/db"
	"medical-matching/routers"

	"github.com/gin-gonic/gin"
)

func initManager() error {
	err := hospital.GetHospitalManager().InitHospitalManager()
	if err != nil {
		return err
	}
	err = hospital.GetSymptomManager().InitSymptomManager()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	r := gin.Default()

	routers.SetupRoutes(r)

	dbManager := db.GetDBManager()
	dbManager.InitDB()

	err := initManager()
	if err != nil {
		log.Panic(err)
	}

	defer dbManager.CloseAll()

	r.Run(":8080")
}
