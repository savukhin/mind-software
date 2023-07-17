package main

import (
	"fmt"
	"log"
	"mind-software/config"
	"strconv"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/nats-io/nats.go"
	"github.com/serialx/hashring"
)

func main() {
	serversInRing := []string{"192.168.0.246:11212",
		"192.168.0.247:11212",
		"192.168.0.248:11212",
		"192.168.0.249:11212",
		"192.168.0.250:11212",
		"192.168.0.251:11212",
		"192.168.0.252:11212"}

	// replicaCount := 3
	ring := hashring.New(serversInRing)
	server, ok := ring.GetNode("my_key")
	fmt.Println(server)
	fmt.Println(ok)

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
	js, err := nc.JetStream()
	if err != nil {
		panic(err)
	}
	js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
	})
	js.UpdateStream(&nats.StreamConfig{
		Name:     "ORDERS",
		MaxBytes: 8,
	})

	for i := 0; ; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Publish", i)
		nc.Publish("foo-nc", []byte("NC Message "+strconv.Itoa(i)))
		js.Publish("ORDERS.scratch", []byte("JS Message "+strconv.Itoa(i)))
	}
}
