package main

import (
	"encoding/base64"
	"encoding/json"
	"hook_pipe/internal/core/settings"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type WebhookMsg struct {
	DeliveryID string            `json:"delivery_id"`
	Method     string            `json:"method"`
	Path       string            `json:"path"`
	Query      string            `json:"query"`
	Headers    map[string]string `json:"headers"`
	BodyBase64 string            `json:"body_base64"`
}

func main() {

	settings.LoadDotEnv()

	settings.LoadEnvs()

	rabbitURL := settings.Settings.RABBIT_URL
	queue := settings.Settings.QUEUE
	exchange := settings.Settings.EXCHANGE
	binding := settings.Settings.BINDING_KEY
	target := settings.Settings.TARGET_BASE

	conn, err := amqp.Dial(rabbitURL)
	must(err)
	defer conn.Close()

	ch, err := conn.Channel()
	must(err)
	defer ch.Close()

	// Asegura que exista el exchange (topic y durable)
	err = ch.ExchangeDeclare(exchange, "topic", true, false, false, false, nil)
	must(err)

	// Declara la cola (durable)
	_, err = ch.QueueDeclare(
		queue,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,
	)
	must(err)

	// Bindea la cola al exchange con la binding key
	err = ch.QueueBind(
		queue,
		binding,  // p.ej. hello.# para todo lo de 'hello'
		exchange, // debe coincidir con el de la Ingress
		false,
		nil,
	)
	must(err)

	// Prefetch 1 para mantener orden (opcional)
	err = ch.Qos(1, 0, false)
	must(err)

	msgs, err := ch.Consume(queue, "", false, false, false, false, nil)
	must(err)

	client := &http.Client{Timeout: 30 * time.Second}

	for d := range msgs {
		var m WebhookMsg
		if err := json.Unmarshal(d.Body, &m); err != nil {
			log.Printf("bad json: %v", err)
			_ = d.Nack(false, false) // DLQ según tu setup
			continue
		}
		body, _ := base64.StdEncoding.DecodeString(m.BodyBase64)

		// Build URL
		url := strings.TrimRight(target, "/") + m.Path
		if m.Query != "" {
			url += "?" + m.Query
		}

		// Sanitize headers
		req, _ := http.NewRequest(m.Method, url, io.NopCloser(strings.NewReader(string(body))))
		for k, v := range m.Headers {
			kL := strings.ToLower(k)
			if kL == "host" || kL == "content-length" || kL == "connection" || kL == "transfer-encoding" {
				continue
			}
			req.Header.Set(k, v)
		}
		req.Header.Set("Idempotency-Key", m.DeliveryID)

		resp, err := client.Do(req)
		if err != nil || resp.StatusCode < 200 || resp.StatusCode >= 300 {
			if err != nil {
				log.Printf("deliver err: %v", err)
			} else {
				log.Printf("deliver status: %d", resp.StatusCode)
			}
			_ = d.Nack(false, false)
			continue
		}
		_ = d.Ack(false)
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
