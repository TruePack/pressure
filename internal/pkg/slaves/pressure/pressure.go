package pressure

import (
	"pressure/internal/pkg/converter"
	"time"

	"github.com/goburrow/modbus"
)

const (
	slaveID  = 16
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
	pressureTransducerClient := Transducer{client: modbusClient}

	return pressureTransducerClient
}

func (t Transducer) GetName() (string, error) {
	registerIndex, err := converter.FromHexToUInt16("0x0000")
	if err != nil {
		return "", err
	}

	name, err := t.client.ReadInputRegisters(registerIndex, 4)
	if err != nil {
		return "", err
	}

	return string(name), nil
}

func (t Transducer) SaveSettings(forever bool) ([]byte, error) {
	value := uint16(0)

	registerIndex, err := converter.FromHexToUInt16("0x000F")
	if err != nil {
		return []byte{}, err
	}

	if forever {
		value = uint16(129)
	}

	saveStatus, err := t.client.WriteSingleRegister(registerIndex, value)
	if err != nil {
		return []byte{}, err
	}

	return saveStatus, nil
}

func (t Transducer) GetPressureResult() (float32, error) {
	registerIndex, err := converter.FromHexToUInt16("0x0013")
	if err != nil {
		return 0, err
	}

	pressure, err := t.client.ReadInputRegisters(registerIndex, 2)
	if err != nil {
		return 0, err
	}

	return converter.FromBytesToFloat32(pressure), nil
}

func (t Transducer) GetPressureUnit() (string, error) {
	registerIndex, err := converter.FromHexToUInt16("0x0015")
	if err != nil {
		return "", err
	}

	pressureUnit, err := t.client.ReadInputRegisters(registerIndex, 1)
	if err != nil {
		return "", err
	}

	return convertToPressureUnit(pressureUnit), nil
}

func (t Transducer) SetPressureUnit(pressureUnit string) (string, error) {
	registerIndex, err := converter.FromHexToUInt16("0x0015")
	if err != nil {
		return "", err
	}

	valueForSlave := convertFromPressureUnitToUInt16(pressureUnit)

	pressureUnitAsBytes, err := t.client.WriteSingleRegister(registerIndex, valueForSlave)
	if err != nil {
		return "", err
	}

	return convertToPressureUnit(pressureUnitAsBytes), nil
}
