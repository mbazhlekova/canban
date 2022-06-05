package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mbazhlekova/canban/config"
	"github.com/mbazhlekova/canban/routes"
)

func main() {
	r := gin.Default()

	routes.ProjectRoute(r)

	config.ConnectDB()

	r.Run("localhost:8081")
}
