package handlers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func Stream(c *fiber.Ctx) error{
	suuid:= c.Params("suuid")
	if suuid==""{
		c.Status(400)
		return nil
	}
	ws:="ws"
	if os.Getenv("ENIRONMENT")=="PRODUCTION" {
		ws="wss"
	}
	w.RoomsLock.Lock()
	if _,ok:=w.Streams[suuid]; ok{
		w.RoomsLock.Unlock()
		return c.Render("stream", fiber.Map{
			"StreamWebsocketAddr": fmt.Sprintf("%s://%s/stream/%s/websocket", ws,c.Hostname(), suuid),
			"ChatWebsocket": fmt.Sprintf("%s://%s/stream/%s/chat/websocket", ws, c.Hostname(), suuid),
			"ViewerWebsocketAddr": fmt.Sprintf("%s://%s/stream/%s/chat/websocket"ws, c.Hostname(),suuid),
			"Type": "stream",
			}, "layouts/main")
	}
	w.RoomsLock.Unlock()
	return c.Render("stream", fiber.Map{
		"NoStream": "true",
		"Leave": "true",
		}, "layouts/main")
}

func StreamWebSocket(c *websocket.Conn){
	suuid:= c.Params("suuid")
	if suuid==""{
		return
	}
	w.RoomsLock.Lock()
	if stream, ok :=w.Stream[uuid]; ok{
		w.StreamConn(c, streas.Peers)
		return
	}
	w.RoomsLock.Unlock()
} 
func StreamViewerWebSocket(c *websocket.Conn){
	suuid:= c.Params("suuid")
	if suuid==""{
		return
	}
	w.RoomsLock.Lock()
	if stream, ok :=w.Stream[uuid]; ok{
		ViewerConn(c, streas.Peers)
		return
	}
	w.RoomsLock.Unlock()
}

func ViewerConn(c *websocket.Conn, p *w.Peers){
	ticker:=time.NewTicker(1*time.Second)
	defer ticker.Stop()
	defer c.Close()
	for {
		select{
		case <-ticker.C:
			w, err:= c.Conn.NextWriter(websocket.TextMessage)
			if err!=nil{
				return
			}
			w.Write([]byte(fmt.Sprintf("%d", len(p.connections))))
		}
	}
}