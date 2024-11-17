package main

import (
	"MedicalMatching/db"
	"MedicalMatching/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.SetupRoutes(r)

	db.InitDB()

	defer db.CloseAll()

	r.Run(":8080")
}
