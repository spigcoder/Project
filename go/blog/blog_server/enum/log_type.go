package enum

type LogType int8

const (
	LoginLogType   LogType = iota + 1 //登录日志
	OperateLogType                    //操作日志
	RuntimeLogType                    //运行日志
)
