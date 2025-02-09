package database

import (
    "context"
    "log"
    "time"
    "go.mongodb.org/mongo-driver/bson"
)

// Define the Alert struct
type Alert struct {
    User          string    `bson:"user"`
    Asset         string    `bson:"asset"`
    PriceThreshold float64  `bson:"priceThreshold"`
    Above         bool      `bson:"above"`
    CreatedAt     time.Time `bson:"createdAt"`
}

// Ensure you call GetCollection from the database package
var alertCollection = GetCollection("alerts-db", "alerts") // Use the GetCollection function here

// CreateAlert inserts a new alert into the MongoDB collection
func CreateAlert(alert Alert) error {
    alert.CreatedAt = time.Now()
    _, err := alertCollection.InsertOne(context.TODO(), alert)
    if err != nil {
        log.Printf("Error inserting alert: %v", err)
        return err
    }
    log.Println("Alert created successfully")
    return nil
}

// GetAlerts fetches all alerts from the MongoDB collection
func GetAlerts() ([]Alert, error) {
    var alerts []Alert
    cursor, err := alertCollection.Find(context.TODO(), bson.M{})
    if err != nil {
        log.Printf("Error fetching alerts: %v", err)
        return nil, err
    }
    defer cursor.Close(context.TODO())

    // Loop through the cursor to decode alerts
    for cursor.Next(context.TODO()) {
        var alert Alert
        if err := cursor.Decode(&alert); err != nil {
            log.Printf("Error decoding alert: %v", err)
            continue
        }
        alerts = append(alerts, alert)
    }
    return alerts, nil
}