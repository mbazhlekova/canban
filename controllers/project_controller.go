package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mbazhlekova/canban/config"
	"github.com/mbazhlekova/canban/models"
	"github.com/mbazhlekova/canban/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var projectCollection *mongo.Collection = config.GetCollection(config.DB, "projects")
var validate = validator.New()

func CreateProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var project models.Project
		defer cancel()

		if err := c.BindJSON(&project); err != nil {
			c.JSON(http.StatusBadRequest, responses.ProjectResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&project); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ProjectResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newProject := models.Project{
			Id:          primitive.NewObjectID(),
			Name:        project.Name,
			Description: project.Description,
			Columns:     project.Columns,
		}

		result, err := projectCollection.InsertOne(ctx, newProject)
		if err != nil {
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.ProjectResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}
		c.JSON(http.StatusCreated, responses.ProjectResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetProject() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("id")
		var project models.Project
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(id)

		err := projectCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&project)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ProjectResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.ProjectResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": project}})
	}
}
