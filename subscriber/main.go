package main

import (
	"fmt"
	"time"

	"github.com/IDarar/hub/pkg/logger"
	nats "github.com/nats-io/nats.go"
)

func main() {
	url := "nats://localhost:4222"
	nc, err := nats.Connect(url)
	if err != nil {
		logger.Info(err)
		return
	}
	defer nc.Close()

	count := 0

	sub, err := nc.Subscribe("events.*", func(msg *nats.Msg) {
		msg.Respond([]byte("123"))
		count++
		logger.Info(msg, string(msg.Data))
	})

	if err != nil {
		logger.Info(err)
		return
	}
	defer sub.Unsubscribe()

	for {
		old := count
		time.Sleep(5 * time.Second)
		if old == count {
			sub.Unsubscribe()
			break
		}
	}
	fmt.Println(count)
}
