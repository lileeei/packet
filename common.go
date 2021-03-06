package packet

import (
	"errors"
	"math"
	"fmt"
)

type Packet struct {
	pos  int32  //偏移量
	Data []byte //数据
}

const (
	INITLEN = 128
)

func (pkt *Packet)Pos() int32 {

	return pkt.pos
}

func (pkt *Packet) Seek(offset int32) bool {
	newPos := pkt.pos + offset
	if len(pkt.Data) < int(newPos) {

		return false
	}

	pkt.pos += offset

	return true
}

func NewPacket() *Packet {
	return &Packet{
		pos:  0,
		Data: make([]byte, 0, INITLEN),
	}
}

//---------------------------------------------------------------------------读取函数
//读取固定数量个byte
func (pkt *Packet) ReadByte(v ...uint32) (bytes []byte, err error) {
	byteNum := append(v, 1)[0]

	if uint32(len(pkt.Data)) < byteNum {
		return nil, errors.New("read byte error")
	}

	pos := pkt.pos
	pkt.pos += int32(byteNum)

	return pkt.Data[pos:pkt.pos], nil
}

func (pkt *Packet) ReadBool() (value bool,err error) {
	bt, err := pkt.ReadByte()
	if err != nil {
		err = errors.New("read bool error")
		return
	}

	return bt[0] == byte(1), nil
}

func (pkt *Packet) ReadInt8() (value int8, err error) {
	ret, reterr := pkt.ReadUint8()
	if reterr == nil {
		value = int8(ret)

		return
	}

	err = errors.New("read int8 error")

	return
}

func (pkt *Packet) ReadUint8() (value uint8, err error) {
	data, err := pkt.ReadByte()
	if err != nil {
		err = errors.New("read uint8 error")
		return
	}

	value = uint8(data[0])

	return
}

func (pkt *Packet) ReadInt16() (value int16, err error) {
	ret, reterr := pkt.ReadUint16()
	if reterr == nil {
		value = int16(ret)

		return
	}

	err = errors.New("read int16 error")

	return
}

func (pkt *Packet) ReadUint16() (value uint16, err error) {
	data, err := pkt.ReadByte(uint32(2))
	if err != nil {
		err = errors.New("read uint16 error")
		return
	}

	value = uint16( data[1]<<8 ) | uint16(data[0])

	return
}

func (pkt *Packet) ReadInt32() (value int32, err error) {
	ret, reterr := pkt.ReadUint32()
	if reterr != nil {
		value = int32(ret)

		return
	}

	err = errors.New("read int32 error")

	return
}

func (pkt *Packet) ReadUint32() (value uint32, err error) {
	data, err := pkt.ReadByte(uint32(4))
	if err != nil {
		err = errors.New("read uint32 error")
		return
	}

	for k, v := range data {
		value |= uint32(v<<uint32((3-k)*8))
	}


	return
}

func (pkt *Packet) ReadInt64() (value int64, err error) {
	ret, reterr := pkt.ReadUint64()
	if reterr == nil {
		value = int64(ret)

		return
	}

	err = errors.New("read int64 error")

	return
}

func (pkt *Packet) ReadUint64() (value uint64, err error) {
	data, err := pkt.ReadByte(uint32(8))
	if err != nil {
		err = errors.New("read uint64 error")
		return
	}

	for k, v := range data {
		value |= uint64(v<<uint32((7-k)*8))
	}

	return
}

func (pkt *Packet) ReadFloat32() (value float32, err error) {
	data, err := pkt.ReadUint32()
	if err != nil {
		err = errors.New("read float32 error")
		return
	}

	value = math.Float32frombits(data)

	return
}

func (pkt *Packet) ReadString(v ...int) (value string, err error) {
	len, reterr := pkt.ReadUint8()
	if reterr != nil {
		err = errors.New("read string len error")
		return
	}

	if len <= 253 {
		ret, reterr:= pkt.ReadByte(uint32(len))

		if reterr != nil {
			err = errors.New("read string error")

			return
		}

		value = string(ret)

		return
	}

	if len == 254 {
		retstrLen, reterr := pkt.ReadUint16()
		if reterr != nil {
			err = errors.New("read string error")
			return
		}

		ret, reterr:= pkt.ReadByte(uint32(retstrLen))
		if reterr != nil {
			err = errors.New("read string error")
			return
		}

		value = string(ret)

		return
	}

	if len == 255 {
		retstrLen,reterr := pkt.ReadUint32()
		if reterr != nil {
			err = errors.New("read string error")
			return
		}

		ret, reterr:= pkt.ReadByte(retstrLen)
		if reterr != nil {
			err = errors.New("read string error")

			return
		}

		value = string(ret)

		return
	}

	return
}

//-----------------------------------------------------------------------写入函数

func (pkt *Packet) WriteBytes(elems []byte) {
	pkt.Data = append(pkt.Data, elems[0:]...)
	//fmt.Printf("write %v bytes(byte)-------------------\n", len(elems))
	pkt.Seek(int32(len(elems)))
}

func (pkt *Packet) WriteBool(elem bool) {
	//fmt.Printf("writed a bool data\n")
	if elem {
		pkt.WriteBytes([]byte{1})
	} else {
		pkt.WriteBytes([]byte{0})
	}
}

func (pkt *Packet) WriteInt8(elem int8) {
	//fmt.Printf("writed a int8 data\n")
	pkt.WriteUint8(uint8(elem))
}

func (pkt *Packet) WriteUint8(elem uint8) {
	pkt.WriteBytes([]byte{elem})
}

func (pkt *Packet) WriteUint16(elem uint16) {
	bt := make([]byte, 2)
	bt[0] = byte(elem >> 8)
	bt[1] = byte(elem)

	pkt.WriteBytes(bt)
}

func (pkt *Packet) WriteInt16(elem int16) {
	//fmt.Printf("writed a int16 data\n")
	pkt.WriteUint16(uint16(elem))
}

func (pkt *Packet) WriteUint32(elem uint32) {
	bt := make([]byte, 4)
	bt[0] = byte(elem >> 24)
	bt[1] = byte(elem >> 16)
	bt[2] = byte(elem >> 8)
	bt[3] = byte(elem)

	pkt.WriteBytes(bt)
}

func (pkt *Packet) WriteInt32(elem int32) {
	fmt.Printf("writed a int32 data\n")
	pkt.WriteUint32(uint32(elem))
}

func (pkt *Packet) WriteUint64(elem uint64) {
	bt := make([]byte, 8)
	for i := 0; i < 8; i++ {
		bt[i] = byte(elem >> (7-1)*8)
	}

	pkt.WriteBytes(bt)
}

func (pkt *Packet) WriteInt64(elem int64) {
	//fmt.Printf("writed a int64 data\n")
	pkt.WriteUint64(uint64(elem))
}

func (pkt *Packet) WriteFloat32(elem float32) {
	//fmt.Printf("writed a float32 data\n")
	pkt.WriteUint32(math.Float32bits(elem))
}

func (pkt *Packet) WriteFloat64(elem float64) {
	//fmt.Printf("writed a float64 data\n")
	pkt.WriteUint64(math.Float64bits(elem))
}

func (pkt *Packet) WriteString(elem string) {
	//fmt.Printf("writed a string data\n")
	strLen := len(elem)
	if strLen <= 253 {
		pkt.WriteUint8(uint8(strLen))
	} else if strLen < 2^16 {
		pkt.WriteUint8(uint8(254))
		pkt.WriteUint16(uint16(strLen))
	} else {
		pkt.WriteUint8(uint8(255))
		pkt.WriteUint32(uint32(strLen))
	}

	pkt.WriteBytes([]byte(elem))
}
