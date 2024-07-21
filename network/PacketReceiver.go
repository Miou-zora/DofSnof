package network

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type PacketSniffer struct {
	Buffer        chan []byte
	MaxBufferSize int32
	Filter        string
	Device        string
}

func (receiver *PacketSniffer) Run() {
	go receiver.sniff()
}

func (receiver *PacketSniffer) sniff() {
	for true {
		handle, err := pcap.OpenLive(receiver.Device, receiver.MaxBufferSize, false, pcap.BlockForever)
		if err != nil {
			log.Fatal(err)
		}
		defer handle.Close()
		err = handle.SetBPFFilter(receiver.Filter)
		if err != nil {
			fmt.Println("Error: ", err)
			log.Fatal(err)
		}
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			receiver.Buffer <- packet.Layer(layers.LayerTypeTCP).LayerPayload()
		}
	}
}
