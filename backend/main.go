package main

import (
	"MedicalMatching/controller/hospital"
	"MedicalMatching/db"
	"MedicalMatching/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.SetupRoutes(r)

	dbManager := db.GetDBManager()
	dbManager.InitDB()

	err := hospital.GetHospitalManager().InitHospitalManager()
	if err != nil {
		log.Panic(err)
	}
	err = hospital.GetSymptomManager().InitSymptomManager()
	if err != nil {
		log.Panic(err)
	}

	defer dbManager.CloseAll()

	r.Run(":8080")
}
