package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"time"

	"awesomeProject/functions"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	uri := os.Getenv("MONGO_URI")
	db := "dev"

	client, err := functions.ConnectToMongoDB(uri)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	fmt.Println("Connected to MongoDB!")
	database := client.Database(db)

	startTime := time.Now()
	seedDataAndHandleError(database, "users", functions.UserCount, functions.GenerateUser, "Seeding user data...")
	seedDataAndHandleError(database, "companies", functions.CompanyCount, functions.GenerateCompany, "Seeding company and return data...")

	fmt.Printf("Seeded data! (%.2f seconds, %d documents)\n",
		time.Since(startTime).Seconds(),
		functions.UserCount+functions.CompanyCount*functions.ReturnCount)
}

func seedDataAndHandleError(database *mongo.Database, collectionName string, count int, generate func() functions.Entity, message string) {
	fmt.Println(message)
	err := functions.SeedData(database, database.Collection(collectionName), count, generate)
	if err != nil {
		panic(err)
	}
}
