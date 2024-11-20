package main

import "zinx/znet"

// Server模块的测试函数

func main() {
	s := znet.NewServer("[zinx V0.1]")
	s.Server()
}
