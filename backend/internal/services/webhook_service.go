package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/DavidAfdal/purchasing-systeam/internal/dto"
)

type webhook struct {
	URL string
}

type WebhookService interface {
	SendWebhook(payload dto.WebhookPayload) error
}

func NewWebhookService(url string) WebhookService {
	return &webhook{url}
}

func (w *webhook) SendWebhook(payload dto.WebhookPayload) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Post(w.URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		log.Printf("Webhook returned status %d\n", resp.StatusCode)
	}

	log.Println("Webhook sent successfully")
	return nil
}
