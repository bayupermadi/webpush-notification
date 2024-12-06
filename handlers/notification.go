package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bayupermadi/webpush-notification/config"

	"github.com/SherClockHolmes/webpush-go"
)

type Subscription struct {
	Endpoint string `json:"endpoint"`
	Keys     struct {
		P256dh string `json:"p256dh"`
		Auth   string `json:"auth"`
	} `json:"keys"`
}

var subscriptions []Subscription

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	var sub Subscription
	err := json.NewDecoder(r.Body).Decode(&sub)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	subscriptions = append(subscriptions, sub)
	w.WriteHeader(http.StatusOK)
}

func SendNotificationHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, sub := range subscriptions {
			resp, err := webpush.SendNotification([]byte("Hello from Go!"), &webpush.Subscription{
				Endpoint: sub.Endpoint,
				Keys: webpush.Keys{
					P256dh: sub.Keys.P256dh,
					Auth:   sub.Keys.Auth,
				},
			}, &webpush.Options{
				Subscriber:      cfg.Subscriber,
				VAPIDPublicKey:  cfg.PublicKey,
				VAPIDPrivateKey: cfg.PrivateKey,
			})

			if err != nil {
				log.Println("Error sending notification:", err)
				continue
			}
			defer resp.Body.Close()
			log.Println("Notification sent:", resp.Status)
		}
		w.WriteHeader(http.StatusOK)
	}
}
