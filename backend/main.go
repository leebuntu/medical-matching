package main

import (
	"log"
	"medical-matching/controller/hospital"
	"medical-matching/db"
	"medical-matching/db/providers"
	"medical-matching/routers"
	"medical-matching/test"

	"github.com/gin-gonic/gin"
)

func injectData() error {
	sm := test.GetSymptomInjection()
	hm := test.GetHospitalInjection()

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
	symptoms, err := providers.GetSymptomProvider().FetchSymptoms()
	if err != nil {
		return err
	}

	err = hospital.GetSymptomManager().ResetSymptomManager(symptoms)
	if err != nil {
		return err
	}

	hospitals, err := providers.GetHospitalProvider().FetchHospitals()
	if err != nil {
		return err
	}

	err = hospital.GetHospitalManager().ResetHospitalManager(hospitals)
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
