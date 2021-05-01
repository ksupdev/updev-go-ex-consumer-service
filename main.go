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

	defer ms.Cleanup()
	ms.Start()
}

func randString() string {
	i := rand.Int()
	return fmt.Sprintf("%d", i)
}
