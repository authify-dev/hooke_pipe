package controllers

import (
	"encoding/base64"
	"encoding/json"
	"hook_pipe/internal/core/settings"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type WebhookMsg struct {
	DeliveryID string            `json:"delivery_id"`
	ReceivedAt time.Time         `json:"received_at"`
	Vendor     string            `json:"vendor"`
	Method     string            `json:"method"`
	Path       string            `json:"path"`
	Query      string            `json:"query"`
	Headers    map[string]string `json:"headers"`
	BodyBase64 string            `json:"body_base64"`
	RemoteIP   string            `json:"remote_ip,omitempty"`
}

func routingKey(vendor, path, method string) string {
	p := strings.TrimLeft(path, "/")
	p = strings.ReplaceAll(p, "/", ".")
	if p == "" {
		p = "root"
	}
	return vendor + "." + p + "." + strings.ToUpper(method)
}

func (c *HooksController) Pipe(ctx *gin.Context) {

	rabbitURL := settings.Settings.RABBIT_URL
	exchange := settings.Settings.EXCHANGE

	// Conexión y canal AMQP (una vez, al arranque)
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("AMQP dial: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("AMQP channel: %v", err)
	}
	defer ch.Close()

	if err := ch.ExchangeDeclare(exchange, "topic", true, false, false, false, nil); err != nil {
		log.Fatalf("declare exchange: %v", err)
	}

	vendor := ctx.Param("vendor")
	rawPath := ctx.Param("path")
	if rawPath == "" {
		rawPath = "/"
	}

	body, err := ctx.GetRawData()
	if err != nil {
		ctx.Status(400)
		return
	}

	headers := make(map[string]string, len(ctx.Request.Header))
	for k, v := range ctx.Request.Header {
		if len(v) > 0 {
			headers[strings.ToLower(k)] = v[0]
		}
	}

	msg := WebhookMsg{
		DeliveryID: time.Now().UTC().Format("20060102T150405.000000000Z07:00"),
		ReceivedAt: time.Now().UTC(),
		Vendor:     vendor,
		Method:     ctx.Request.Method,
		Path:       rawPath,
		Query:      ctx.Request.URL.RawQuery,
		Headers:    headers,
		BodyBase64: base64.StdEncoding.EncodeToString(body),
		RemoteIP:   ctx.ClientIP(),
	}

	data, _ := json.Marshal(msg)
	rk := routingKey(vendor, rawPath, ctx.Request.Method)

	pub := amqp.Publishing{
		ContentType:  "application/json",
		Body:         data,
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now().UTC(),
		Type:         "webhook.event",
		MessageId:    msg.DeliveryID,
	}

	// Publica y responde
	if err := ch.Publish(exchange, rk, false, false, pub); err != nil {
		log.Printf("publish error rk=%s: %v", rk, err)
		ctx.Status(500)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
