package multimeter

import (
	"encoding/binary"
	"pressure/internal/pkg/converter"
	"time"

	"github.com/goburrow/modbus"
)

const (
	slaveID  = 1
	parity   = "N"
	baudrate = 9600
	COMport  = "COM3"
)

type Transducer struct {
	client modbus.Client
}

func GetClient() Transducer {
	handler := modbus.NewRTUClientHandler(COMport)
	handler.BaudRate = baudrate
	handler.Parity = parity
	handler.SlaveId = slaveID
	handler.Timeout = 2 * time.Second

	modbusClient := modbus.NewClient(handler)
	frequenceTransducerClient := Transducer{client: modbusClient}

	return frequenceTransducerClient
}

func (t Transducer) GetAddressAndSpeed() (uint16, uint16, error) {
	registerIndex, err := converter.FromHexToUInt16("0x0001")
	if err != nil {
		return 0, 0, err
	}

	addressAndSpeed, err := t.client.ReadInputRegisters(registerIndex, 1)
	if err != nil {
		return 0, 0, err
	}

	address, speed := splitBytesFromDualRegister(addressAndSpeed)

	addressAsUInt := binary.BigEndian.Uint16(address)
	speedAsUInt := binary.BigEndian.Uint16(speed)

	return addressAsUInt, speedAsUInt, nil
}

func (t Transducer) GetDotAndTypeForPower() (uint16, uint16, error) {
	registerIndex, err := converter.FromHexToUInt16("0x0018")
	if err != nil {
		return 0, 0, err
	}

	addressAndSpeed, err := t.client.ReadInputRegisters(registerIndex, 1)
	if err != nil {
		return 0, 0, err
	}

	address, speed := splitBytesFromDualRegister(addressAndSpeed)

	addressAsUInt := binary.BigEndian.Uint16(address)
	speedAsUInt := binary.BigEndian.Uint16(speed)

	return addressAsUInt, speedAsUInt, nil
}

func splitBytesFromDualRegister(value []byte) ([]byte, []byte) {
	var address []byte

	var speed []byte

	switch len(value) {
	case 3:
		address = value[0:1]
		speed = value[2:]
	case 4:
		address = value[0:2]
		speed = value[3:]
	}

	return address, speed
}
