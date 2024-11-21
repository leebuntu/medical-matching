package main

import (
	"log"
	"medical-matching/controller/hospital"
	"medical-matching/db"
	"medical-matching/routers"

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
