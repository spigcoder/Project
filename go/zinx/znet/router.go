package znet

import "zinx/ziface"

// BaseRouter类为一切实现了Router子类的父类，BaseRouter实现了IRouter的3个接口
// 但是BaseRouter的方法都是空，这就让有的Router不需要PreHandle的也可以实例化
type BaseRouter struct {}

func(br *BaseRouter)PreHandle(req ziface.IRequest) {}
func(br *BaseRouter)Handle(req ziface.IRequest) {}
func(br *BaseRouter)PostHandle(req ziface.IRequest) {}