package converter

import (
	"encoding/binary"
	"math"
	"strconv"
)

const (
	HexBitSize    = 16
	HexBaseWith0x = 0
)

func FromHexToUInt16(hex string) (uint16, error) {
	convertedValue, err := strconv.ParseUint(hex, HexBaseWith0x, HexBitSize)
	if err != nil {
		return 0, err
	}

	valueAsUInt := uint16(convertedValue)

	return valueAsUInt, nil
}

func FromBytesToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	float := math.Float32frombits(bits)

	return float
}
