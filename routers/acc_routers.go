// Package routers 路由
package routers

import (
	"g2git.hdyouxi.com/xiandi/game-socket-server/servers/websocket"
)

// WebsocketInit Websocket 路由
func WebsocketInit() {
	websocket.Register("login", websocket.LoginController)
	websocket.Register("heartbeat", websocket.HeartbeatController)
	websocket.Register("ping", websocket.PingController)
}
