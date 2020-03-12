package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Database *mongo.Database
var Context context.Context

func getDatabaseConnection() string {
	databaseConn := os.Getenv("DATABASE_CONNECTION")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")

	databaseConn = strings.Replace(databaseConn, "<username>", databaseUsername, 1)
	databaseConn = strings.Replace(databaseConn, "<password>", databasePassword, 1)

	return databaseConn
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env", err.Error())
	}
}

func init() {
	loadEnv()

	fmt.Println("GET CONNECTION")
	databaseConn := getDatabaseConnection()
	databaseName := os.Getenv("DATABASE")

	Context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Client, err := mongo.Connect(Context, options.Client().ApplyURI(
		databaseConn,
	))

	if err != nil {
		log.Fatal("Error connecting to database...", err.Error())
	}
	Database = Client.Database(databaseName)
}
