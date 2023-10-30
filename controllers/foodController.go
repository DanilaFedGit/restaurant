package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/DanilaFedGit/restaurant/database"
	"github.com/DanilaFedGit/restaurant/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")
var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		foodid := c.Param("food_id")
		var food models.Food
		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodid}).Decode(&food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error with find food"})
		}
		c.JSON(http.StatusOK, food)
	}
}
func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var food models.Food
		var menu models.Menu
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := validate.Struct(food)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "menu wasn't found"})
			return
		}
		food.Created_at = time.Parse(time.RFC3339, time.Now()).Format(time.RFC3339)
		food.Update_at = time.Parse(time.RFC3339, time.Now()).Format(time.RFC3339)
		food.ID = primitive.ObjectID()

	}
}
func round(num float64) int {

}
func tofixed(num float64, precusion int) float64 {

}
