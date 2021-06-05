package main

import (
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
	sub, err := nc.SubscribeSync("events.*")
	if err != nil {
		logger.Info(err)
		return
	}

	for {

		msg, err := sub.NextMsg(3 * time.Second)
		if err != nil {
			logger.Info(err)
			return
		}
		go logger.Info(string(msg.Data))

	}

}
