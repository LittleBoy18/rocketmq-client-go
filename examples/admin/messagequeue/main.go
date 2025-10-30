package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	topic := "test"
	resolver := primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})

	a, err := admin.NewAdmin(
		admin.WithResolver(resolver),
		admin.WithCredentials(primitive.Credentials{
			AccessKey: "console",
			SecretKey: "Hc@Cloud01",
		}),
	)
	if err != nil {
		log.Fatalf("new admin error: %v", err)
	}
	defer func() {
		_ = a.Close()
	}()

	mqs, err := a.FetchPublishMessageQueues(context.Background(), topic)
	if err != nil {
		log.Fatalf("fetch message queues error: %v", err)
	}
	if len(mqs) == 0 {
		log.Fatalf("no message queue found for topic: %s", topic)
	}

	for _, mq := range mqs {
		maxOffset, err := a.QueryMessageQueueMaxOffset(mq)
		if err != nil {
			log.Printf("query max offset error for %s[%s:%d]: %v", mq.Topic, mq.BrokerName, mq.QueueId, err)
			continue
		}
		minOffset, err := a.QueryMessageQueueMinOffset(mq)
		if err != nil {
			log.Printf("query min offset error for %s[%s:%d]: %v", mq.Topic, mq.BrokerName, mq.QueueId, err)
			continue
		}
		fmt.Printf("Topic=%s, Broker=%s, QueueId=%d, MinOffset=%d, MaxOffset=%d\n", mq.Topic, mq.BrokerName, mq.QueueId, minOffset, maxOffset)
	}
}
