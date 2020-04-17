package src

import (
	"github.com/google/gopacket"
)


func export(packet gopacket.Packet) {
	srcIP := "nil"
	dstIP := "nil"
	if packet.NetworkLayer != nil && packet.NetworkLayer() != nil {
		srcIP = packet.NetworkLayer().NetworkFlow().Src().String()
		dstIP = packet.NetworkLayer().NetworkFlow().Dst().String()

	}
	srcPort := "nil"
	dstPort := "nil"
	protocol := "nil"
	if packet.TransportLayer != nil && packet.TransportLayer() != nil {
		srcPort = packet.TransportLayer().TransportFlow().Src().String()
		dstPort = packet.TransportLayer().TransportFlow().Dst().String()
		protocol = packet.TransportLayer().LayerType().String()
	}


	packetLabels := []string{
		srcIP,
		dstIP,
		srcPort,
		dstPort,
		protocol}

	countPacket(packetLabels)
}