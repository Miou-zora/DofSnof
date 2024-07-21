package main

import (
	"fmt"
	"log"
	game "sniffsniff/game/network"
	"sniffsniff/network"
	"sniffsniff/utils"
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
		case raw_data := <-receiver.Buffer:
			if len(raw_data) == 0 {
				continue
			}
			buffer = append(buffer, raw_data...)
			header := game.HeaderFromByte(buffer)
			if header.IsValid() {
				fmt.Println("Message: ", game.ID_TO_MESSAGE_NAMES[int(header.Id)])
			} else {
				fmt.Println("Invalid message: ", header.Id)
			}
			size := 2 + int(header.LenType) + int(header.DataLen)
			if size > len(buffer) {
				fmt.Print("Packet is not complete, waiting for more data...")
				continue
			}
			data := utils.Buffer{Data: buffer[(2 + header.LenType):size], Pos: 0}
			buffer = buffer[size:]
			if game.ID_TO_MESSAGE[header.Id] != nil {
				message := game.ID_TO_MESSAGE[header.Id]()
				message.Deserialize(&data)
				fmt.Println("Message: ", message)
			}
		default:
			continue
		}
	}
}
