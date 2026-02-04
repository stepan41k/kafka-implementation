package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "my-topic",
		Balancer: &kafka.LeastBytes{},
	}

	defer writer.Close()

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("Message %d", i)
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("Key-%d", i)),
				Value: []byte(msg),
			},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}
		fmt.Println("Sent:", msg)
		time.Sleep(1 * time.Second)
	}
}