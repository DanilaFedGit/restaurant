package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
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
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}
		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))
		matchStage := bson.D{{"$match", bson.D{}}}
		groupStage := bson.D{{"$group", bson.D{{"_id", bson.D{{"_id", "null"}}}, {"total_sum", bson.D{{"$sum", 1}}}}}}
		projectStage := bson.D{{"$project",bson.D{{"_id",0},{"total_count",1},{"food_items",bson.D{{"$slice",[]interface{}{"$data",startIndex,recordPerPage}}}}}}}
		result, err:=foodCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage,
		})
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}
		var allFoods []bson.M
		err = result.All(ctx,&allFoods)
		if err!=nil{
			log.Fatal(err)
			return
		}
		c.JSON(http.StatusOK,gin.H{"data":&allFoods})
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
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var food models.Food
		var menu models.Menu
		foodid:=c.Param("food_id")
		err := c.BindJSON(&food)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}
	
		var updateObj primitive.D
		if food.Name != "" {
			updateObj = append(updateObj, bson.E{"name", food.Name})
		}
		if food.Price != "" {
			updateObj = append(updateObj, bson.E{"category", menu.Category})
		}
		if food.Food_image!=nil{

		}
		if food.Menu_id!=nil{

		}
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"update_time", menu.Updated_at})
		upsert := true
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		food.Update_at,_ = time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		result, err := menuCollection.UpdateOne(ctx,
			menuid,
			bson.D{
				{"$set", updateObj}},
			&opt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "update wasn't save"})
			return
		}
		c.JSON(http.StatusOK, result)
	}

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
		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()
		var num = tofixed(*food.Price, 2)
		food.Price = &num
		result, err := foodCollection.InsertOne(ctx, food)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "food wasn't created"})
			return
		}
		c.JSON(http.StatusOK, result)

	}
}

func round(num float64) int {

}
func tofixed(num float64, precusion int) float64 {

}
