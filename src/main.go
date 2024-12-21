package main

import (
	"context"
	"crypto/tls"

	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	opensearch "github.com/opensearch-project/opensearch-go/v4"
)

func postgresInit() {
	conn, err := pgx.Connect(context.Background(), "postgres://user:password@localhost:5432/mydatabase")
	if err != nil {
		fmt.Println("Unable to connect to database:", err)
		return
	}
	defer conn.Close(context.Background())
	fmt.Println("Connected to PostgreSQL!")
}

func mongoInit() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Unable to connect to MongoDB:", err)
		return
	}
	defer client.Disconnect(context.Background())

	fmt.Println("Connected to MongoDB!")
}

func neo4jInit() {
	driver, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("username", "password", ""))
	if err != nil {
		fmt.Println("Unable to connect to Neo4j:", err)
		return
	}
	defer driver.Close()

	fmt.Println("Connected to Neo4j!")
}

func redisInit() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // нет пароля
		DB:       0,  // использовать базу по умолчанию
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Unable to connect to Redis:", err)
		return
	}

	fmt.Println("Connected to Redis!")
}

func opensearchInit() {
	username := "admin" // Replace with your username
	password := "Opensearch1_"

	_, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{"https://localhost:9200"},
		Username:  username, // For testing only. Don't store credentials in code.
		Password:  password,
	})

	if err != nil {
		fmt.Println("Not connected")
		return
	}

	fmt.Println("Connected to Opensearch")

}

func main() {
	postgresInit()
	mongoInit()
	neo4jInit()
	redisInit()
	opensearchInit()
}
