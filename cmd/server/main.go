package main

import (
	"log"

	"github.com/GuaiZai233/Larvar/internal/auth"
	"github.com/GuaiZai233/Larvar/internal/db"
	"github.com/GuaiZai233/Larvar/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.GetDB()
	if err := database.AutoMigrate(&auth.User{}); err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	r := gin.Default()
	router.SetupRouter(r)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
