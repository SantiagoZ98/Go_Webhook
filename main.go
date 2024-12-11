package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Estructura de los datos del webhook
type WebhookPayload struct {
	Message string `json:"message"`
}

// Handler para recibir el webhook
func webhookHandler(w http.ResponseWriter, r *http.Request) {
	// Solo aceptamos solicitudes POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decodifica el cuerpo del mensaje JSON
	var payload WebhookPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Responde con un mensaje de confirmación
	fmt.Fprintf(w, "Webhook recibido con mensaje: %s", payload.Message)
}

func main() {
	// Configura el endpoint para el webhook
	http.HandleFunc("/webhook", webhookHandler)

	// Inicia el servidor HTTP
	log.Println("Iniciando servidor en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
