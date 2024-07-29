package network

import (
	"fmt"
	"log"

	"github.com/google/gopacket/pcap"
)

type Device struct {
	Name        string
	Description string
}

func getListOfDevices() ([]Device, error) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return nil, err
	}
	return toDevices(devices), nil
}

func toDevices(deviceInterfaces []pcap.Interface) []Device {
	devices := make([]Device, len(deviceInterfaces))
	for index, device := range deviceInterfaces {
		devices[index] = Device{Name: device.Name, Description: device.Description}
	}
	return devices
}

func AskForDevice() (Device, error) {
	devices, err := getListOfDevices()
	if err != nil {
		return Device{}, err
	}
	for i, device := range devices {
		println(i, ": ", device.Description)
	}
	var deviceIndex int = -1
	for deviceIndex < 0 || deviceIndex >= len(devices) {
		println("Please select a device: ")
		_, err := fmt.Scanf("%d", &deviceIndex)
		if err != nil {
			log.Fatal(err)
		}
	}
	return devices[deviceIndex], nil
}
