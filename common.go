package packet

import (
	"strconv"
)

type PlayLoad struct {
	pos int32   	//偏移量
	Data []byte	//数据
}

const (
	INITLEN = 128
)

func NewPlayLoad() *PlayLoad {
	return &PlayLoad{
		pos: 0,
		Data: make([]byte, 0, INITLEN),
	}
}


//---------------------------------------------------------------------------读取函数
//读取固定数量个byte
func (playload *PlayLoad)ReadByte(v... int) []byte {
	byteNum := append(v, 1)[0]

	if len(playload.Data) < byteNum {
		return nil
	}
	pos := playload.pos
	playload.pos += int32(byteNum)

	return playload.Data[pos:playload.pos]
}

func (playload *PlayLoad)ReadBool() bool {
	bt := playload.ReadByte()
	//if bt == nil {
	//	return nil
	//}

	return bt[0] == byte(1)
}

func (playload *PlayLoad)ReadInt8() int8 {
	data := playload.ReadByte()
	//if Data == nil{
	//	return Data
	//}

	return int8(data[0])
}

func (playload *PlayLoad)ReadUint8() uint8 {
	data := playload.ReadByte()
	//if Data == nil{
	//	return Data
	//}

	return uint8(data[0])
}

func (playload *PlayLoad)ReadInt16() int16{
	bt := playload.ReadByte(2)

	value, _ := strconv.ParseInt(string(bt), 10, 16)
	//if err != nil {
	//	return nil
	//}
	return int16(value)
}


func (playload *PlayLoad)ReadUint16() uint16{
	bt := playload.ReadByte(2)

	value, _ := strconv.ParseUint(string(bt), 10, 16)
	//if err != nil {
	//	return nil
	//}
	return uint16(value)
}


func (playload *PlayLoad)ReadInt32() int32{
	bt := playload.ReadByte(4)

	value, _ := strconv.ParseInt(string(bt), 10, 32)
	//if err != nil {
	//	return nil
	//}

	return int32(value)
}


func (playload *PlayLoad)ReadUint32() uint32{
	bt := playload.ReadByte(4)

	value, _ := strconv.ParseUint(string(bt), 10, 32)
	//if err != nil {
	//	return nil
	//}

	return uint32(value)
}

func (playload *PlayLoad)ReadInt64() int64{
	bt := playload.ReadByte(8)

	value, _ := strconv.ParseInt(string(bt), 10, 64)
	//if err != nil {
	//	return nil
	//}

	return int64(value)
}

func (playload *PlayLoad)ReadUint64() uint64{
	bt := playload.ReadByte(8)

	value, _ := strconv.ParseUint(string(bt), 10, 64)
	//if err != nil {
	//	return nil
	//}

	return uint64(value)
}

func (playload *PlayLoad)ReadFloat32() float32 {
	bt := playload.ReadByte(4)

	value, _ := strconv.ParseFloat(string(bt), 32)
	//if err != nil {
	//	return
	//}

	return float32(value)
}

func (playload *PlayLoad)ReadString(v... int) string{
	leftLen := len(playload.Data) - int(playload.pos)
	num := append(v, leftLen)[0]

	bt := playload.ReadByte(num)
	//if bt != nil {
	//	return string(bt)
	//}

	return string(bt)
}




//-----------------------------------------------------------------------写入函数

func (playload *PlayLoad)WriteBytes(elems... byte) {
	var index int = 0

	//for playload.pos < int32(cap(playload.Data)) && index < len(elems) {
	//	playload.Data[playload.pos] = elems[index]
	//	playload.pos++
	//	index++
	//}
	//
	//if index < len(elems) {
		playload.Data = append(playload.Data, elems[index:]...)
	//}
}

func (playload *PlayLoad)WriteBool(elem bool) {
	if elem {
		playload.WriteBytes(byte(1))
	} else {
		playload.WriteBytes(byte(0))
	}
}

func (playload *PlayLoad)WriteInt8(elem int8)  {
	playload.WriteBytes(byte(elem))
}

func (playload *PlayLoad)WriteUint8(elem uint8) {
	playload.WriteBytes(byte(elem))
}

func (playload *PlayLoad)WriteInt16(elem int16) {
	playload.WriteBytes([]byte(string(elem))...)
}


func (playload *PlayLoad)WriteUint16(elem int16) {
	playload.WriteBytes([]byte(string(elem))...)
}


func (playload *PlayLoad)WriteInt32(elem int32) {
	playload.WriteBytes([]byte(string(elem))...)
}


func (playload *PlayLoad)WriteUint32(elem uint32) {
	playload.WriteBytes([]byte(string(elem))...)
}

func (playload *PlayLoad)WriteInt64(elem int64) {
	playload.WriteBytes([]byte(string(elem))...)
}

func (playload *PlayLoad)WriteUint64(elem uint64){
	playload.WriteBytes([]byte(string(elem))...)
}

func (playload *PlayLoad)WriteFloat32(elem float32) {
	s := strconv.FormatFloat(float64(elem), 'f', -1, 32)
	playload.WriteBytes([]byte(s)...)
}

func (playload *PlayLoad)WriteFloat64(elem float64) {
	s := strconv.FormatFloat(elem, 'f', -1, 64)
	playload.WriteBytes([]byte(s)...)
}

func (playload *PlayLoad)WriteString(elem string) {
	playload.WriteBytes([]byte(elem)...)
}




