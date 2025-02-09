package database

import (
    "context"
    "log"
    "time"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "os"
)

var Client *mongo.Client

func Connect() error {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Read values from environment variables
    uri := os.Getenv("MONGODB_URI")
    database := os.Getenv("MONGODB_DATABASE")
    collection := os.Getenv("MONGODB_COLLECTION")

    // Check if any values are missing
    if uri == "" || database == "" || collection == "" {
        log.Fatalf("Missing required environment variables!")
    }

    // Set up MongoDB client options
    clientOptions := options.Client().ApplyURI(uri)

    // Create new MongoDB client
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        log.Printf("Failed to create MongoDB client: %v", err)
        return err
    }

    // Connect to MongoDB
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Printf("Failed to connect to MongoDB: %v", err)
        return err
    }

    // Assign client to the global variable
    Client = client
    log.Println("Connected to MongoDB!")

    // Log database and collection info
    log.Printf("Using database: %s, collection: %s", database, collection)

    return nil
}

// GetCollection returns the MongoDB collection from the environment variables
// func GetCollection() *mongo.Collection {
//     databaseName := os.Getenv("MONGODB_DATABASE") // Get MongoDB database name from the environment
//     collectionName := os.Getenv("MONGODB_COLLECTION") // Get MongoDB collection name from the environment

//     // Access the collection
//     return Client.Database(databaseName).Collection(collectionName)
// }

// GetCollection accepts database and collection names as arguments
func GetCollection(databaseName, collectionName string) *mongo.Collection {
    // Access the collection using passed arguments
    return Client.Database(databaseName).Collection(collectionName)
}