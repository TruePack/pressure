package pressure

const (
	Pa      = "Па"
	KPa     = "кПа"
	VPI     = "% ВПИ"
	MMH2O   = "мм.вод.ст"
	MH2O    = "м.вод.ст"
	MBar    = "мБар"
	Bar     = "Бар"
	Psi     = "psi"
	Nothing = ""
)

func convertToPressureUnit(pressure []byte) string {
	switch pressure[0] {
	case 0:
		return Pa
	case 1:
		return KPa
	case 2:
		return VPI
	case 3:
		return MMH2O
	case 4:
		return MH2O
	case 5:
		return MBar
	case 6:
		return Bar
	case 7:
		return Psi
	}

	return Nothing
}

func convertFromPressureUnitToUInt16(pressure string) uint16 {
	switch pressure {
	case Pa:
		return 0
	case KPa:
		return 1
	case VPI:
		return 2
	case MMH2O:
		return 3
	case MH2O:
		return 4
	case MBar:
		return 5
	case Bar:
		return 6
	case Psi:
		return 7
	}

	return 0
}
