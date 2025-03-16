package enum

type LogLevel int8

const (
	InfoLogLevel LogLevel = iota + 1
	WarnLogLevel
	ErrorLogLevel
)

func (l LogLevel) String() string {
	switch l {
	case InfoLogLevel:
		return "info"
	case WarnLogLevel:
		return "warn"
	case ErrorLogLevel:
		return "err"
	}
	return ""
}
