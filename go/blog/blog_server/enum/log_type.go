package enum

type LogType int8

const (
	LoginLogType   LogType = iota
	OperateLogType         //操作日志
	RuntimeLogType         //运行日志
)
