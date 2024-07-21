package network

import (
	"log"
	"github.com/google/gopacket/pcap"
)

type Device struct {
	Name string
	Description string
}

func GetListOfDevices() []Device {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatalf("Error finding devices: %v", err)
	}
	return toDevices(devices)
}

func toDevices(deviceInterfaces []pcap.Interface) []Device {
	var devices []Device
	for _, device := range deviceInterfaces {
		devices = append(devices, Device{Name: device.Name, Description: device.Description})
	}
	return devices
}