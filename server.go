package main

import (
	"ZINX_PRO/zinx/ziface"
	"ZINX_PRO/zinx/znet"
	"fmt"
	"mmo_game/api"
	"mmo_game/core"
)

//当客户端建立连接的时候的hook函数
func OnConnecionAdd(conn ziface.IConnection) {
	//创建一个玩家
	player := core.NewPlayer(conn)
	//同步当前的PlayerID给客户端， 走MsgID:1 消息
	player.SyncPid()

	//同步当前玩家的初始化坐标信息给客户端，走MsgID:200消息
	player.BroadCastStartPosition()

	//将当前新上线玩家添加到worldManager中
	core.WorldMgrObj.AddPlayer(player)

	//将该连接绑定属性Pid
	conn.SetProperty("pid", player.Pid)

	fmt.Println("=====> Player pidId = ", player.Pid, " arrived ====")
}

func main() {
	//创建服务器句柄
	s := znet.NewServer("MMO Game -ZINX")

	//注册客户端连接建立和丢失函数
	s.SetOnConnStart(OnConnecionAdd)

	//注册一些路由业务
	s.AddRouter(2, &api.WorldChatApi{})

	//启动服务
	s.Serve()
}
