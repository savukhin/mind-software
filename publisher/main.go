package main

import (
	"fmt"
	"log"
	"math/rand"
	"publisher/config"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/nats-io/nats.go"
)

type NatsLogMsg struct {
	ObjId uint64 `json:"id"`
	Log   string `json:"log"`
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	cfg := &config.Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	nc, err := nats.Connect(cfg.NatsUrl)
	if err != nil {
		panic(err)
	}
	defer nc.Close()
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		panic(err)
	}
	defer ec.Close()

	sleepDuration := time.Duration(int(time.Minute) / cfg.QueriesMin)

	for i := 0; ; i++ {
		time.Sleep(sleepDuration)
		// nc.Publish("foo-nc", []byte("NC Message "+strconv.Itoa(i)))

		msg := &NatsLogMsg{
			ObjId: uint64(rand.Intn(3)),
			Log:   randSeq(rand.Intn(100)),
		}
		fmt.Printf("Publish i = %d; msg log = '%s' objId = %d\n", i, msg.Log, msg.ObjId)
		ec.Publish(cfg.NatsLogsTopic, msg)
	}
}
