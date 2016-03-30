package dock

type ModuleType int

const (
	Unknown ModuleType = iota
	Joystick
	Slider
	Touch
	Motion
	Light
	Colour
	Dial
	Barometer
	Number
	Motor
	Rainbow
	Matrix
)

var modules map[string]ModuleType

func init() {
	modules = make(map[string]ModuleType)
	modules["slider"] = Slider
	modules["touch"] = Touch
	modules["motion"] = Motion
	modules["light"] = Light
	modules["colour"] = Colour
	modules["dial"] = Dial
	modules["barometer"] = Barometer
	modules["number"] = Number
	modules["motor"] = Motor
	modules["rainbow"] = Rainbow
	modules["joystick"] = Joystick
	modules["matrix"] = Matrix
}

func FromString(name string) ModuleType {
	if m, ok := modules[name]; ok {
		return m
	}
	return Unknown
}

func (m ModuleType) String() string {
	switch m {
	case Slider:
		return "slider"
	case Touch:
		return "touch"
	case Motion:
		return "motion"
	case Light:
		return "light"
	case Colour:
		return "colour"
	case Dial:
		return "dial"
	case Barometer:
		return "barometer"
	case Number:
		return "number"
	case Motor:
		return "motor"
	case Rainbow:
		return "rainbow"
	case Joystick:
		return "joystick"
	case Matrix:
		return "matrix"
	}
	return "unknown"
}

type Module struct {
	ModuleType
	*Dock
	port int
}

func (m *Module) ProcessEvent(ev Event) {
	if ev.ModuleType == m.ModuleType {
		if ev.EventType == Connected {
			m.port = ev.Port
		}
		if ev.EventType == Disconnected {
			m.port = -1
		}
	}
}

func (m *Module) Connected() bool {
	return m.port != -1
}

func (m *Module) Port() int {
	return m.port
}

func (m *Module) Set(params ...int) error {
	return m.Dock.SetModuleData(m.port, m.ModuleType, params...)
}
