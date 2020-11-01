package converter_test

import (
	"pressure/internal/pkg/converter"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromHexToUInt16(t *testing.T) {
	hex1 := "0x0000"
	hex2 := "0x000F"
	hex3 := "0x0010"
	hex4 := "0x0111"
	hex5 := "0x0074"

	expected1 := uint16(0x0)
	expected2 := uint16(0xf)
	expected3 := uint16(0x10)
	expected4 := uint16(0x111)
	expected5 := uint16(0x74)

	actual1, err1 := converter.FromHexToUInt16(hex1)
	actual2, err2 := converter.FromHexToUInt16(hex2)
	actual3, err3 := converter.FromHexToUInt16(hex3)
	actual4, err4 := converter.FromHexToUInt16(hex4)
	actual5, err5 := converter.FromHexToUInt16(hex5)

	assert.NoError(t, err1)
	assert.NoError(t, err2)
	assert.NoError(t, err3)
	assert.NoError(t, err4)
	assert.NoError(t, err5)

	assert.Equal(t, expected1, actual1)
	assert.Equal(t, expected2, actual2)
	assert.Equal(t, expected3, actual3)
	assert.Equal(t, expected4, actual4)
	assert.Equal(t, expected5, actual5)
}
