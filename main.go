package main

import (
	"fmt"
	game "sniffsniff/game/network"
	"sniffsniff/network"
	"log"
)

func askForDevice() string {
	deviceNames := network.GetListOfDevices()
	for i, device := range deviceNames {
		println(i, ": ", device.Description)
	}
	var deviceIndex int = -1
	for deviceIndex < 0 || deviceIndex >= len(deviceNames) {
		println("Please select a device: ")
		_, err := fmt.Scanf("%d", &deviceIndex)
		if err != nil {
			log.Fatal(err)
		}
	}
	return deviceNames[deviceIndex].Name
}

func main() {
	deviceName := askForDevice()

	buffer := make([]byte, 0)

	receiver := network.PacketSniffer{
		Buffer:        make(chan []byte, 4096),
		MaxBufferSize: 1600,
		Filter:        "tcp src port 5555",
		Device:        deviceName,
	}
	receiver.Run()
	for {
		select {
		case data := <-receiver.Buffer:
			buffer = append(buffer, data...)
			header := game.HeaderFromByte(buffer)
			if header.IsValid() {
				fmt.Println("Message: ", game.ALL_MESSAGE[int(header.Id)])
			} else {
				fmt.Println("Invalid message: ", header.Id)
			}
            // TODO: Check is buffer is big enough to hold the header + message
			buffer = make([]byte, 0) // It should only erase the size of the header + message
		default:
			continue
		}
	}
}
