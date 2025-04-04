package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rezacloner1372/hesabketab/internal/api"
)

func main() {
	router := gin.Default()
	api.RegisterRoutes(router)
	const port = ":8080"
	log.Printf("Server is running on port %s", port)
}
