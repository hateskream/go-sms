package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"space-management-system/app"
	"space-management-system/services/db/db"
	"space-management-system/services/reservations"
	"space-management-system/services/spaces"

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

	connConfig, err := pgx.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatalf("Unable to parse connection config: %v", err)
	}
	conn, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(ctx)
	log.Printf("Connected to database: %s", connConfig.Database)

	queries := db.New(conn)
	spaceManager := &spaces.SpacesManager{}
	ReservationsManager := &reservations.ReservationsManager{}

	app := app.InitializeApp()
	app.Storage = queries
	app.SetSpaces(spaceManager)
	app.SetReservations(ReservationsManager)

	router := initRoutes()
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}
	server.ListenAndServe()
}
