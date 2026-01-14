package main

import (
	"learing_go/simple_crud/internal/app/api/db"
	"learing_go/simple_crud/internal/app/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run("localhost:8080")
}
