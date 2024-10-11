package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"Backend_Room_Booking/database"
	"Backend_Room_Booking/gmail"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	gmailService := gmail.InitializeGmailClient()

	database.ConnectDatabase()

	go gmail.ProcessIncomingEmails(gmailService)

	log.Println("Backend running...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Room Booking API!"))
	})

	log.Printf("Starting HTTP server on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}

	select {}
}
