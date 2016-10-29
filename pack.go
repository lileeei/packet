package packet

import (
	"reflect"
)


//
func Pack(writer *Packet, tbl interface{}) []byte {
	if writer == nil {
		writer = NewPacket()
	}

	s := reflect.ValueOf(tbl)
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		pack(writer, f)
	}

	return writer.Data
}

func pack(writer *Packet, value reflect.Value) []byte {
	switch value.Type().Kind() {
	case reflect.Struct:
		Pack(writer, value.Interface())

	case reflect.Map:
		pack_map(writer, value)

	case reflect.Array, reflect.Slice:
		pack_arrayorslice(writer, value)

	case reflect.Ptr, reflect.Interface:
		pack(writer, value.Elem())

	default:
		if is_primitive(value) {
			pack_primitive(writer, value)
		}
	}

	return writer.Data
}

func is_primitive(e reflect.Value) bool {

	switch e.Kind() {
	case reflect.Bool,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Float32,
		reflect.Float64,
		reflect.String:

		return true
	}

	return false
}

func pack_primitive(writer *Packet, value reflect.Value) {

	switch value.Kind() {
	case reflect.Bool:
		writer.WriteBool(value.Bool())

	case reflect.Uint8:
		writer.WriteUint8(uint8(value.Uint()))
	case reflect.Uint16:
		writer.WriteUint16(uint16(value.Uint()))
	case reflect.Uint32:
		writer.WriteUint32(uint32(value.Uint()))
	case reflect.Uint64:
		writer.WriteUint64(uint64(value.Uint()))

	case reflect.Int8:
		writer.WriteInt8(int8(value.Int()))
	case reflect.Int16:
		writer.WriteInt16(int16(value.Int()))
	case reflect.Int32:
		writer.WriteInt32(int32(value.Int()))
	case reflect.Int64:
		writer.WriteInt64(int64(value.Int()))

	case reflect.Float32:
		writer.WriteFloat32(float32(value.Float()))
	case reflect.Float64:
		writer.WriteFloat64(float64(value.Float()))

	case reflect.String:
		writer.WriteString(string(value.String()))
	}
}

func pack_arrayorslice(writer *Packet, value reflect.Value) {
	len := value.Len()
	if len <= 253 {
		writer.WriteUint8(uint8(len))
	} else if len < 2^16 {
		writer.WriteUint8(uint8(254))
		writer.WriteUint16(uint16(len))
	} else {
		writer.WriteUint8(uint8(255))
		writer.WriteUint32(uint32(len))
	}

	for i := 0; i < len; i++ {
		v := value.Index(i)
		pack(writer, v)
	}
}

func pack_map(writer *Packet, value reflect.Value) {
	keySlice := value.MapKeys()

	for i := 0; i < value.Len(); i++ {
		key := keySlice[i]
		pack(writer, key)
		_value := value.MapIndex(key)
		pack(writer, _value)
	}
}
