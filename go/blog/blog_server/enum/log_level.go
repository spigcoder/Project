package enum

type LogLevel int8

const (
    InfoLogLevel LogLevel = iota
    WarnLogLevel
    ErrorLogLevel
)