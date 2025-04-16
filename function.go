package main

import (
	"fmt"
	"os/exec"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func isOnline(mac string) bool {
	cmd := exec.Command("arp", "-a")
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.Contains(string(output), mac)
}

func actionTopicWol(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	// Function to run
	WakeOnLAN(COMPUTER_MAC_ADDRESS)
}

func actionTopicTestConnect(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	if isOnline(COMPUTER_MAC_ADDRESS) {
		publishData("deviceControl/device-monitor", "online")
	} else {
		publishData("deviceControl/device-monitor", "offline")
	}
}
