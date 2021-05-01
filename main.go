package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ms := NewMicroservice()

	servers := "localhost:9094"
	topic := "when-citizen-has-registered-" + randString()
	groupID := "validation-consumer"
	timeout := time.Duration(-1)

	ms.Consume(servers, topic, groupID, timeout, func(ctx IContext) error {
		msg := ctx.ReadInput()
		ctx.Log(msg)
		return nil
	})

	prod := NewProducer(servers, ms)
	go func() {
		for i := 0; i < 10; i++ {
			prod.SendMessage(topic, "", map[string]interface{}{"message_id": i})
			time.Sleep(time.Second)
		}

		// Exit program
		ms.Stop()
	}()

	defer ms.Cleanup()
	ms.Start()
}

func randString() string {
	i := rand.Int()
	return fmt.Sprintf("%d", i)
}
