package utils

import(
	"encoding/json"
	"io/ioutil"
	"zinx/ziface"
)

// 这里主要是存储一些Zinx框架的全局函数，供其他模块进行使用，
// 一些参数可以让用户通过zinx.json进行配置

type GlobalObj struct {
	TcpServer 	ziface.IServer
	Host		string
	TcpPort		int
	Name		string
	Version		string

	MaxPacketSize uint32
	MaxConn		int
}

var GlobalObject *GlobalObj

// 读取用户的配置文件
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	//将json数据解析到struct当中去, 这里还要传递指针是为了让GolbalObject指向一个新的空间
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name: "ZinxServerApp",
		Version: "V0.4",
		TcpPort: 7777,
		Host: "0.0.0.0",
		MaxConn: 12000,
		MaxPacketSize: 4096,
	}
	GlobalObject.Reload()
}