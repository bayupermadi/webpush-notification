package vapid

import (
	"fmt"
	"log"

	"github.com/SherClockHolmes/webpush-go"
)

func GenerateVAPIDKeys() {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		log.Fatalf("Error generating VAPID keys: %v", err)
	}

	fmt.Println("Public Key:", publicKey)
	fmt.Println("Private Key:", privateKey)
}
