package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func init() {
	// Load .env
	if err := godotenv.Load(".env"); err != nil {
		log.Println("warning: .env not loaded:", err)
	}

	mongoURI := os.Getenv("MONGOSTRING")
	if mongoURI == "" {
		log.Fatal("MONGOSTRING kosong. Pastikan .env atau env sudah diset.")
	}

	dbName := os.Getenv("MONGODB_NAME")
	if dbName == "" {
		dbName = "tes_db" // default
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Mongo connect error:", err)
	}

	// Pastikan benar-benar terkoneksi
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Mongo ping error:", err)
	}

	MongoDB = client.Database(dbName)
	log.Println("✅ MongoDB connected:", dbName)
}
