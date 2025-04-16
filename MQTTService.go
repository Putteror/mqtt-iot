package main

import (
	"fmt"
	"log"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

func publishData(topic string, message string) {

	opts := MQTT.NewClientOptions()
	opts.AddBroker(BROKER_ADDRESS)
	opts.SetClientID(uuid.New().String())

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	payload := message
	token := client.Publish(topic, 1, false, payload)
	token.Wait()

	time.Sleep(1 * time.Second)
	client.Disconnect(250)
}

func subscribeData() {

	opts := MQTT.NewClientOptions()
	opts.AddBroker(BROKER_ADDRESS)
	client_id := uuid.New().String()
	opts.SetClientID(client_id)
	opts.OnConnect = func(c MQTT.Client) {
		fmt.Println("Connected to MQTT broker")
		if token := c.Subscribe("deviceControl/wol", 1, actionTopicWol); token.Wait() && token.Error() != nil {
			log.Fatal(token.Error())
		}
		if token := c.Subscribe("deviceControl/test-connect", 1, actionTopicTestConnect); token.Wait() && token.Error() != nil {
			log.Fatal(token.Error())
		}
	}

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	for {
		time.Sleep(1 * time.Second)
	}
}
