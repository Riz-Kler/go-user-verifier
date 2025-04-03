package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/segmentio/kafka-go"
)

type VerificationRequest struct {
    UserID string `json:"userID"`
}

type VerificationResult struct {
    UserID string `json:"userID"`
    Status string `json:"status"`
}

func main() {
    broker := os.Getenv("KAFKA_BROKER")
    if broker == "" {
        broker = "localhost:9092"
    }

    requestTopic := "user.verification.request"
    resultTopic := "user.verification.result"

    reader := kafka.NewReader(kafka.ReaderConfig{
        Brokers:   []string{broker},
        Topic:     requestTopic,
        GroupID:   "verifier-group",
        MinBytes:  1e3,
        MaxBytes:  1e6,
    })
    defer reader.Close()

    writer := &kafka.Writer{
        Addr:     kafka.TCP(broker),
        Topic:    resultTopic,
        Balancer: &kafka.LeastBytes{},
    }
    defer writer.Close()

    fmt.Println("✅ Verifier service is running and listening for Kafka events...")

    for {
        msg, err := reader.ReadMessage(context.Background())
        if err != nil {
            log.Printf("❌ Failed to read message: %v", err)
            continue
        }

        var req VerificationRequest
        if err := json.Unmarshal(msg.Value, &req); err != nil {
            log.Printf("❌ Invalid message format: %v", err)
            continue
        }

        fmt.Printf("🔍 Verifying user: %s\n", req.UserID)
        time.Sleep(2 * time.Second) // Simulate verification

        result := VerificationResult{
            UserID: req.UserID,
            Status: "verified",
        }

        value, _ := json.Marshal(result)
        err = writer.WriteMessages(context.Background(), kafka.Message{
            Key:   []byte(req.UserID),
            Value: value,
        })
        if err != nil {
            log.Printf("❌ Failed to write result: %v", err)
        } else {
            log.Printf("✅ User %s verified and result sent to Kafka.", req.User
