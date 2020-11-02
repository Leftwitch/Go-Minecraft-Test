package packets

import (
	"bufio"
	"bytes"
	"errors"
)

// Packet Is the base interface for every Packet
type Packet interface {
	GetPacketID() uint16
	ReadPacket(Buffer *bufio.ReadWriter)
	WritePacket(Buffer *bytes.Buffer)
}

// ReadVarInt Reads a VarInt from the Buffer
func ReadVarInt(Buffer *bufio.ReadWriter) (int32, error) {
	var n int32
	for i := 0; ; i++ { //读数据前的长度标记
		sec, err := Buffer.ReadByte()
		if err != nil {
			return 0, err
		}

		n |= int32(sec&0x7F) << int32(7*i)

		if i >= 5 {
			return 0, errors.New("VarInt is too big")
		} else if sec&0x80 == 0 {
			break
		}
	}

	return n, nil
}

//WriteVarInt c
func WriteVarInt(Buffer *bytes.Buffer, num int32) {
	for {
		b := num & 0x7F
		num >>= 7
		if num != 0 {
			b |= 0x80
		}
		Buffer.WriteByte(byte(b))
		if num == 0 {
			break
		}
	}
	return
}

//ReadNBytes read N bytes from bytes.Reader
func ReadNBytes(Buffer *bufio.ReadWriter, n int32) (bs []byte, err error) {
	bs = make([]byte, n)
	for i := int32(0); i < n; i++ {
		bs[i], err = Buffer.ReadByte()
		if err != nil {
			return
		}
	}
	return
}

//WriteString c
func WriteString(Buffer *bytes.Buffer, str string) {
	ByteString := []byte(str)
	WriteVarInt(Buffer, int32(len(ByteString)))
	Buffer.Write(ByteString)
}

//ReadString reads a string from the buffer
func ReadString(Buffer *bufio.ReadWriter) (string, error) {

	StringLegnth, error := ReadVarInt(Buffer)

	if error != nil {
		//Error Handling
		return "", error
	}

	if bytes, error := ReadNBytes(Buffer, StringLegnth); error == nil {
		return string(bytes), nil
	}

	return "", nil

}

//ReadUnsignedShort Reads a Short
func ReadUnsignedShort(Buffer *bufio.ReadWriter) (uint16, error) {
	bs, err := ReadNBytes(Buffer, 2)
	if err != nil {
		return 0, err
	}
	return uint16(int16(bs[0])<<8 | int16(bs[1])), nil
}
