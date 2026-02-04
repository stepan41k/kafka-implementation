package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "my-topic",
		GroupID: "my-group", // Группа нужна для отслеживания прочитанных сообщений
	})

	defer reader.Close()

	fmt.Println("Start consuming...")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read message:", err)
		}
		fmt.Printf("Received: %s = %s\n", string(m.Key), string(m.Value))
	}
}