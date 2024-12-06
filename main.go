package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bayupermadi/webpush-notification/config"
	"github.com/bayupermadi/webpush-notification/vapid"

	"github.com/bayupermadi/webpush-notification/handlers"
)

func main() {
	// Define flags
	generateVapid := flag.Bool("generate-vapid", false, "Generate VAPID keys for push notifications")
	configPath := flag.String("config", "./config.yaml", "Path to the configuration file")

	// Parse flags
	flag.Parse()

	if *generateVapid {
		// Generate VAPID keys and exit
		vapid.GenerateVAPIDKeys()
		os.Exit(0)
	}

	// Load config
	cfg := config.LoadConfig(*configPath)
	fmt.Println("Config loaded successfully.")

	// Setup HTTP Handlers
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/subscribe", handlers.SubscribeHandler)
	http.HandleFunc("/sendNotification", handlers.SendNotificationHandler(cfg))

	// Start Server
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
