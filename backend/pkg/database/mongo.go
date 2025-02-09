package database

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/joho/godotenv"
    "os"
)

var client *mongo.Client
var alertCollection *mongo.Collection

func Connect() error {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    uri := os.Getenv("MONGODB_URI")
    databaseName := os.Getenv("MONGODB_DATABASE")
    collectionName := os.Getenv("MONGODB_COLLECTION")

    // Set client options
    clientOptions := options.Client().ApplyURI(uri)
    client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        return err
    }

    // Ping the database
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        return err
    }

    alertCollection = client.Database(databaseName).Collection(collectionName)
    log.Println("Connected to MongoDB!")
    return nil
}

// GetCollection returns a collection reference
func GetCollection(database, collection string) *mongo.Collection {
    return client.Database(database).Collection(collection)
}

// Example CreateAlert function for reference
func CreateAlert(user, asset string, priceThreshold float64, above bool) error {
    alert := bson.M{
        "user":           user,
        "asset":          asset,
        "priceThreshold": priceThreshold,
        "above":          above,
        "createdAt":      time.Now().Unix(), // Convert to Unix timestamp
    }

    _, err := alertCollection.InsertOne(context.TODO(), alert)
    if err != nil {
        log.Printf("Error inserting alert: %v", err)
        return err
    }

    log.Println("Alert created successfully")
    return nil
}
