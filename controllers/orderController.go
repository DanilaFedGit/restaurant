package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/DanilaFedGit/restaurant/database"
	"github.com/DanilaFedGit/restaurant/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var allOrders []bson.M
		err = result.All(ctx, &allOrders)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders)

	}

}
func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		orderid := c.Param("order_id")
		var order models.Order
		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderid}).Decode(&order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": order})
	}
}
func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var order models.Order
		var table models.Table
		err := c.BindJSON(&order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		err = validate.Struct(order)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
				return
			}
		}
		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()
		result, err := orderCollection.InsertOne(ctx, order)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var order models.Order
		var table models.Table
		var updateObj primitive.D
		orderid := c.Param("order_id")
		err := c.BindJSON(&order)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error})
				return
			}
			updateObj = append(updateObj, bson.E{"table_id", order.Table_id})
		}
		order.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"update_at", order.Update_at})
		upsert := true
		filter := bson.M{"order_id": orderid}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}
		result, err := orderCollection.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}}, &opt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
func OrderItemOrderCreator(order models.Order) string {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()
	orderCollection.InsertOne(ctx, order)
	return order.Order_id
}
