package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"backend/internal/models"
	"backend/pkg/database"
)

func CreateAlert(c *gin.Context) {
	var alert models.Alert
	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	alert.CreatedAt = time.Now().Unix()
	collection := database.GetCollection("alerts-db", "alerts")
	_, err := collection.InsertOne(context.TODO(), alert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save alert"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alert created successfully!"})
}

func GetAlerts(c *gin.Context) {
	collection := database.GetCollection("alerts-db", "alerts")
	cursor, err := collection.Find(context.TODO(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch alerts"})
		return
	}
	defer cursor.Close(context.TODO())

	var alerts []models.Alert
	if err := cursor.All(context.TODO(), &alerts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse alerts"})
		return
	}

	c.JSON(http.StatusOK, alerts)
}