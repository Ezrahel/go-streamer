package webrtc

import (
	"log"
	"github.com/gofiber/websocket"
	"github.com/pion/webrtc/v3"
	"sync"
)

func RoomConn(c *websocket.Conn, p *Peers) {
	var config webrtc.Configuration
	peerConnection,err:=webrtc.NewPeerConnection(config)
	if err!=nil{
		log.Print(err)
		return
	}
	newPeer := PeerConnectionState{
		PeerConnection: peerCOnnection,
		WebSocket:      &ThreadSaferWriter{},
		Conn:           c,
		Mutex:          sync.Mutex{}, 
	}
}