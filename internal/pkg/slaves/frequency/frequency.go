package frequency

import (
	"pressure/internal/pkg/converter"
	"time"

	"github.com/goburrow/modbus"
)

const (
	slaveID  = 11
	parity   = "N"
	baudrate = 4800
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

func (t Transducer) GetFrequency() (float32, error) {
	registerIndex, err := converter.FromHexToUInt16("0x0001")
	if err != nil {
		return 0, err
	}

	frequency, err := t.client.ReadInputRegisters(registerIndex, 1)
	if err != nil {
		return 0, err
	}

	return converter.FromBytesToFloat32(frequency), nil
}

func (t Transducer) SetFrequencyInput() (float32, error) {
	registerIndex, err := converter.FromHexToUInt16("0x0065")
	if err != nil {
		return 0, err
	}

	valueForModBus := uint16(5)
	frequency, err := t.client.WriteSingleRegister(registerIndex, valueForModBus)
	if err != nil {
		return 0, err
	}

	return converter.FromBytesToFloat32(frequency), nil
}
