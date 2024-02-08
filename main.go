package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"space-management-system/app"
	"space-management-system/handlers"
	"space-management-system/hardware"
	"space-management-system/services/db/db"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var port string

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("env loaded")
	port = os.Getenv("PORT")
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get database connection parameters from environment variables
	connConfig, err := pgx.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatalf("Unable to parse connection config: %v", err)
	}

	// Establish connection to the database
	ctx := context.Background()
	conn, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(ctx)

	// Log the database name
	log.Printf("Connected to database: %s", connConfig.Database)

	queries := db.New(conn)
	hardwareMethods := &hardware.Methods{Storage: queries}
	apiHandlers := &handlers.Handlers{Storage: queries, Hardware: hardwareMethods}
	app := &app.App{
		Storage:  queries,
		Hardware: hardwareMethods,
		Handlers: apiHandlers,
	}

	router := initRoutes(app)
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}
	server.ListenAndServe()
}
