package main

import (
	"MedicalMatching/constants"
	"MedicalMatching/db"
	"MedicalMatching/db/user"
	"MedicalMatching/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.SetupRoutes(r)

	db.InitDB()
	userDB, err := db.GetDBManager().GetDB(constants.UserDB)
	if err != nil {
		log.Fatal(err)
	}
	err = user.NewPriorityInjection(userDB).InjectPriority(5, 26)
	if err != nil {
		log.Fatal(err)
	}

	defer db.CloseAll()

	r.Run(":8080")
}
