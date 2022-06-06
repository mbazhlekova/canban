package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mbazhlekova/canban/controllers"
)

func ProjectRoute(router *gin.Engine) {
	router.POST("/project", controllers.CreateProject())
	router.GET("/project/:id", controllers.GetProject())
	router.PUT("/project/:id", controllers.UpdateProject())
	router.DELETE("/project/:id", controllers.DeleteProject())
	router.GET("/projects", controllers.GetAllProjects())
}
