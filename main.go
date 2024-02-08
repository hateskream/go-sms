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

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=localhost port=5432 user=root password=root dbname=root")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	log.Println("Connected db")

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
