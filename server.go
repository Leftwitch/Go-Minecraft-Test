package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/Leftwitch/RestTest/packets"
)

func main() {

	PORT := ":25565"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {

		Buffer := bufio.NewReadWriter(bufio.NewReader(c), bufio.NewWriter(c))

		PacketLength, _ := ReadVarInt(Buffer)
		fmt.Printf("PACKET LENGTH: %d", PacketLength)

		PacketID, _ := ReadVarInt(Buffer)

		var packet packets.Packet

		if PacketID == 0 {
			packet = &packets.PacketHandshake{}
			packet.ReadPacket(Buffer)
			fmt.Println(packet)

			var TempBuffer = new(bytes.Buffer)
			packets.WriteVarInt(TempBuffer, 0) //Packet id
			var PingPacket = new(packets.PacketServerList)
			PingPacket.WritePacket(TempBuffer) //Packet data

			var FinalizedPacket = new(bytes.Buffer)
			packets.WriteVarInt(FinalizedPacket, int32(len(TempBuffer.Bytes()))) //Write packet size
			FinalizedPacket.Write(TempBuffer.Bytes())                            //Write packet data with id

			Buffer.Write(FinalizedPacket.Bytes()) //Send to connection
			Buffer.Flush()                        //Flush
		}

		//c.Close()
		break

	}
}

//ReadVarInt a VarInt
func ReadVarInt(r *bufio.ReadWriter) (uint32, error) {
	var n uint32
	for i := 0; ; i++ { //读数据前的长度标记
		sec, err := r.ReadByte()
		if err != nil {
			return 0, err
		}

		n |= uint32(sec&0x7F) << uint32(7*i)

		if i >= 5 {
			return 0, errors.New("VarInt is too big")
		} else if sec&0x80 == 0 {
			break
		}
	}

	return n, nil
}
