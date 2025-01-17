package webrtc

import (
	"log"
	"sync"

	"github.com/ezrahel/go-streamer/pkg/chat"
	"github.com/ezrahel/go-streamer/pkg/webrtc"
	"github.com/pion/webrtc/v3"
)

type Peers struct {
	ListLock    sync.RWMutex
	Connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

type Rooms struct {
	Peers *Peers
	Hub   *chat.Hub
}

type PeerConnectionState struct {
	PeerConnection *webrtc.PeerConection
	websocket *ThreadSaferWriter
}

type ThreadSaferWriter struct {
	Conn *websocket.Conn
	Mutex sync.Mutex
}
func (t *ThreadSaferWriter) WriteJSON (v interface()) error {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	return t.Conn.WriteJSON(v)
}

func (p *Peers) AddTrack (t *webrtc.TrackRemote) *webrtc.TrackLocalStaticRTP{
	p.ListLock.Lock()
	defer func (){
		p.ListLock.Unlock()
		p.SignalPeerConnection()
	}()
	trackLocal, err := webrtc.NewTrackLocalStaticRTP(t.Codec().RTPCodeCapability, t.ID(). t.StreamID())
		if err!= nil{
			log.Println(err.Error())
			return
		}
		p.TrackLocals[t.ID()] = trackLocal
		return trackLocal
}
func (p *Peers) RemoveTrack (t *webrtc.TrackLocalStaticRTP){
	p.ListLock.Lock()
	defer func(){
		p.ListLock.Unlock()
		p.SignalPeerConnection()
	}()
	delete(p.TrackLocals, t.ID)
}
func (p *Peers) SignalPeerConnection (){
	p.ListLock.Lock()
	defer func(){
		p.ListLock.Unlock()
		p.DispatchKeyFrame()
	}()
	attemptSync := func()(tryAgain bool){
		for i := range p.Connections{
			if p.Connections[i].PeerConnection.ConnectionState()== webrtc.PeerConnectionStateClosed{
				p.Connections= append(p.Connections[:i], p.Connections[1+i:]...)
				log.PrintLn("a", p.Connections)
				return true
			}
			existingSenders:= map[string]bool{}
			for _,sender := range p.Connections[i].PeerConnection.GetSenders(){
				if sender.Track()==nil{
					continue
				}
				existingSenders[senders.Track().ID()]=true
				if _,ok:= p.TrackLocals[sender.Track().ID()]; !ok{
					if err:= p.Connections[i].PeerConnection.RemoveTrack(sender); err!=nil{
						return true
					}
				}
			}
			for _, receiver := range p.Connections[i].PeerConnection.GetReceivers(){
				if receiver.Track()== nil{

					continue
				}
				existingSenders[receiver.Track().ID()]=true
			}
			for trackID := range p.TrackLocals{
				if _, ok := existingSenders [trackID]; !ok{
					if _, err:= p.Connections[i].PeerConnection.AddTrack(p.TrackLocals[trackID]; err!=nil{
						return true
					})
				}
			}
		}
	}
}
func (p *Peers) DispatchKeyFrame() {

}

type websocketMessage struct {
	Event string `json:event"`
	Data string
}