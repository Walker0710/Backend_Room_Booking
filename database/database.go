package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func ConnectDatabase() {
	var err error
	postgresURI := os.Getenv("POSTGRES_URI")
	db, err = pgx.Connect(context.Background(), postgresURI)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	log.Println("Connected to PostgreSQL")
}

func UpdateRoomStatus(roomID string, status string) error {
	_, err := db.Exec(context.Background(), "UPDATE rooms SET status = $1 WHERE id = $2", status, roomID)
	return err
}