package main

import (
	"Test_Task_0/internal/models"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"os"
)

var err error

func main() {
	vv, err := os.ReadFile("./cmd/sender_client/test.json")
	if err != nil {
		logrus.Fatal(err)
	}
	var orders []models.Order
	err = json.Unmarshal([]byte(vv), &orders)
	if err != nil {
		logrus.Fatal(err)
	}

	sc, err := stan.Connect("wbl0-cluster", "sender_client-1", stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		logrus.Fatal(err)
	}
	for idx, order := range orders {
		o, err := json.Marshal(order)
		if err != nil {
			if err != nil {
				logrus.Error(err)
			}
		}
		err = sc.Publish("wb", o)
		if err != nil {
			logrus.Error(err)
		}
		logrus.Printf("message [%d] send succesfull,  uuid:[%s]", idx, order.OrderUid)
	}
}
