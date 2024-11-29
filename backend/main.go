package main

import (
	"log"
	"medical-matching/db"
	"medical-matching/db/hospital"
	"medical-matching/db/sets"
	"medical-matching/routers"

	"github.com/gin-gonic/gin"
)

func injectData() error {
	sm := sets.GetSymptomInjection()
	hm := sets.GetHospitalInjection()

	err := sm.InjectSymptoms()
	if err != nil {
		return err
	}

	err = hm.InjectHospital()
	if err != nil {
		return err
	}
	return nil
}

func resetManagers() error {
	err := hospital.GetHospitalManager().ResetHospitalManager()
	if err != nil {
		return err
	}

	err = hospital.GetSymptomManager().ResetSymptomManager()
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

	defer dbManager.CloseAll()

	err := injectData()
	if err != nil {
		log.Panic(err)
	}

	err = resetManagers()
	if err != nil {
		log.Panic(err)
	}

	r.Run(":8080")
}
