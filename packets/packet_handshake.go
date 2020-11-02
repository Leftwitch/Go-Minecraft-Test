package packets

import (
	"bufio"
	"bytes"
	"fmt"
)

//PacketHandshake c
type PacketHandshake struct {
	ProtocolVersion int32
	ServerAddress   string
	ServerPort      uint16
	NextState       uint8
}

//GetPacketID c
func (packet *PacketHandshake) GetPacketID() uint16 {
	return 0
}

//ReadPacket c
func (packet *PacketHandshake) ReadPacket(Buffer *bufio.ReadWriter) {
	fmt.Printf("Reading STATUS PACKET. \n")

	if ProtocolVersion, error := ReadVarInt(Buffer); error == nil {
		packet.ProtocolVersion = ProtocolVersion
		fmt.Printf("Read Protocol Version %d!\n", ProtocolVersion)
	}

	if ServerAddress, error := ReadString(Buffer); error == nil {
		packet.ServerAddress = ServerAddress
		fmt.Printf("Read Server Address! %s\n", ServerAddress)
	}

	if ServerPort, error := ReadUnsignedShort(Buffer); error == nil {
		packet.ServerPort = ServerPort
		fmt.Printf("Read Server Port! %d\n", ServerPort)
	}

	if NextState, error := ReadVarInt(Buffer); error == nil {
		packet.NextState = uint8(NextState)
		fmt.Printf("Read next state! %d\n", NextState)
	}
	fmt.Printf("Reading Done!")

}

//WritePacket C
func (packet *PacketHandshake) WritePacket(Buffer *bytes.Buffer) {
	fmt.Printf("Writing STATUS PACKET")
}
