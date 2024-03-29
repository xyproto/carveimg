package truetype

import (
	"encoding/binary"
	"fmt"
)

// Code generated by bin-parser-gen. DO NOT EDIT

func (out *VarAxis) mustParse(data []byte) {
	_ = data[19] // early bound checking
	out.Tag = Tag(binary.BigEndian.Uint32(data[0:4]))
	out.Minimum = Float1616FromUint(binary.BigEndian.Uint32(data[4:8]))
	out.Default = Float1616FromUint(binary.BigEndian.Uint32(data[8:12]))
	out.Maximum = Float1616FromUint(binary.BigEndian.Uint32(data[12:16]))
	out.flags = uint16(binary.BigEndian.Uint16(data[16:18]))
	out.strid = NameID(binary.BigEndian.Uint16(data[18:20]))
}

func parseVarAxis(data []byte) (VarAxis, error) {
	var out VarAxis
	if L := len(data); L < 20 {
		return VarAxis{}, fmt.Errorf("EOF: expected length: 20, got %d", L)
	}
	out.mustParse(data)
	return out, nil
}
func (item VarAxis) writeTo(data []byte) {
	_ = data[19] // early bound checking
	binary.BigEndian.PutUint32(data, uint32(item.Tag))
	binary.BigEndian.PutUint32(data[4:], uint32(Float1616ToUint(item.Minimum)))
	binary.BigEndian.PutUint32(data[8:], uint32(Float1616ToUint(item.Default)))
	binary.BigEndian.PutUint32(data[12:], uint32(Float1616ToUint(item.Maximum)))
	binary.BigEndian.PutUint16(data[16:], uint16(item.flags))
	binary.BigEndian.PutUint16(data[18:], uint16(item.strid))
}

func (item VarAxis) appendTo(data []byte) []byte {
	L := len(data)
	data = append(data, make([]byte, 20)...)
	dst := data[L:]
	item.writeTo(dst)
	return data
}

func (out *fvarHeader) mustParse(data []byte) {
	_ = data[15] // early bound checking
	out.majorVersion = uint16(binary.BigEndian.Uint16(data[0:2]))
	out.minorVersion = uint16(binary.BigEndian.Uint16(data[2:4]))
	out.axesArrayOffset = uint16(binary.BigEndian.Uint16(data[4:6]))
	out.reserved = uint16(binary.BigEndian.Uint16(data[6:8]))
	out.axisCount = uint16(binary.BigEndian.Uint16(data[8:10]))
	out.axisSize = uint16(binary.BigEndian.Uint16(data[10:12]))
	out.instanceCount = uint16(binary.BigEndian.Uint16(data[12:14]))
	out.instanceSize = uint16(binary.BigEndian.Uint16(data[14:16]))
}

func parseFvarHeader(data []byte) (fvarHeader, error) {
	var out fvarHeader
	if L := len(data); L < 16 {
		return fvarHeader{}, fmt.Errorf("EOF: expected length: 16, got %d", L)
	}
	out.mustParse(data)
	return out, nil
}
func (item fvarHeader) writeTo(data []byte) {
	_ = data[15] // early bound checking
	binary.BigEndian.PutUint16(data, uint16(item.majorVersion))
	binary.BigEndian.PutUint16(data[2:], uint16(item.minorVersion))
	binary.BigEndian.PutUint16(data[4:], uint16(item.axesArrayOffset))
	binary.BigEndian.PutUint16(data[6:], uint16(item.reserved))
	binary.BigEndian.PutUint16(data[8:], uint16(item.axisCount))
	binary.BigEndian.PutUint16(data[10:], uint16(item.axisSize))
	binary.BigEndian.PutUint16(data[12:], uint16(item.instanceCount))
	binary.BigEndian.PutUint16(data[14:], uint16(item.instanceSize))
}

func (item fvarHeader) appendTo(data []byte) []byte {
	L := len(data)
	data = append(data, make([]byte, 16)...)
	dst := data[L:]
	item.writeTo(dst)
	return data
}
