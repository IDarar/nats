package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/IDarar/hub/pkg/logger"
	nats "github.com/nats-io/nats.go"
)

var rg = rand.New(rand.NewSource(time.Now().Unix()))

func main() {

	url := "nats://localhost:4222"

	//opts := nats.Options{}

	nc, err := nats.Connect(url)
	if err != nil {
		logger.Info(err)
		return
	}

	defer nc.Close()
	i := 0
	for ; i < 1e3; i++ {
		s := fmt.Sprintf("Msg: %v data: %v\n", i, rg.Intn(10000000))

		/*err := nc.Publish("events.123", []byte(s))
		if err != nil {
			logger.Info(err)
			return
		}*/
		msg, err := nc.Request("events.13", []byte(s), 1*time.Millisecond)
		if err != nil {
			logger.Info(err)
			return
		}
		logger.Info(string(msg.Data))
	}
	fmt.Println(i)
}
