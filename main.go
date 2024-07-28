package main

import (
	"fmt"
	"log"
	game "sniffsniff/game/network"
	"sniffsniff/network"
	"sniffsniff/utils"
)

const (
	MaxBufferSize = 4096
	DefaultFilter = "tcp src port 5555"
)

func askForDevice() network.Device {
	devices := network.GetListOfDevices()
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
	return devices[deviceIndex]
}

func main() {
	device := askForDevice()

	buffer := make([]byte, 0)

	receiver := network.PacketSniffer{
		Buffer: make(chan []byte, MaxBufferSize),
		Filter: DefaultFilter,
		Device: device,
	}
	receiver.Run()
	for {
		select {
		case raw_data := <-receiver.Buffer:
			if len(raw_data) == 0 {
				continue
			}
			buffer = append(buffer, raw_data...)
			for len(buffer) > 2 {
				header := game.HeaderFromByte(buffer)
				if header.IsValid() {
					fmt.Println("Message: ", game.ID_TO_MESSAGE_NAMES[int(header.Id)])
				} else {
					fmt.Println("Invalid message: ", header.Id)
					buffer = buffer[:0]
					continue
				}
				size := header.GetSize()
				if size > len(buffer) {
					fmt.Print("Packet is not complete, waiting for more data...")
					continue
				}
				data := utils.Buffer{Data: buffer[(2 + header.LenType):size], Pos: 0}
				buffer = buffer[size:]
				if game.ID_TO_MESSAGE[header.Id] != nil {
					message := game.ID_TO_MESSAGE[header.Id]()
					err := message.Deserialize(&data)
					if err != nil {
						fmt.Println("Error deserializing message: ", err)
					} else {
						fmt.Println("Message: ", message)
					}
				}
			}
		default:
			continue
		}
	}
}
