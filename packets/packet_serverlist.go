package packets

import (
	"bufio"
	"bytes"
	"fmt"
)

//PacketServerList c
type PacketServerList struct {
	json string
}

//GetPacketID c
func (packet *PacketServerList) GetPacketID() uint16 {
	return 0
}

//ReadPacket c
func (packet *PacketServerList) ReadPacket(Buffer *bufio.ReadWriter) {

}

//WritePacket C
func (packet *PacketServerList) WritePacket(Buffer *bytes.Buffer) {
	fmt.Printf("Writing STATUS PACKET")
	WriteString(Buffer, "{ \"version\": { \"name\": \"1.8.7\", \"protocol\": 47 }, \"players\": { \"max\": 100, \"online\": 5, \"sample\": [ { \"name\": \"thinkofdeath\", \"id\": \"4566e69f-c907-48ee-8d71-d7ba5aa00d20\" } ] }, \"description\": { \"text\": \"Hello world\" }, \"favicon\": \"data:image/png;base64,<data>\" }")

}
