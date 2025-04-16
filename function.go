package main

import (
	"fmt"
	"os/exec"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func macAddressIsOnline(mac string) bool {
	cmd := exec.Command("arp", "-a")
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.Contains(string(output), mac)
}

func pingDevice(host string) bool {
	// สำหรับ Linux/Unix: -c คือจำนวนครั้ง, สำหรับ Windows ใช้ -n
	cmd := exec.Command("ping", "-c", "1", host)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Ping error:", err)
		return false
	}

	return strings.Contains(string(output), "1 received")
}

func actionTopicWol(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	// Function to run
	WakeOnLAN(COMPUTER_MAC_ADDRESS)
}

func actionTopicTestConnect(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())

	if pingDevice(COMPUTER_IP_ADDRESS) {
		publishData("deviceControl/device-monitor", "online")
	} else {
		publishData("deviceControl/device-monitor", "offline")
	}
}
