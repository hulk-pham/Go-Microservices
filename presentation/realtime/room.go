package realtime

import (
	"fmt"
	"hulk/go-webservice/domain/entities"

	"github.com/dustin/go-broadcast"
)

type Listener struct {
	RoomId string
	Chan   chan interface{}
}

var RoomManager *Manager

type Manager struct {
	roomChannels map[string]broadcast.Broadcaster
	open         chan *Listener
	close        chan *Listener
	delete       chan string
	messages     chan *entities.Message
}

func InitRoomManager() *Manager {
	RoomManager = &Manager{
		roomChannels: make(map[string]broadcast.Broadcaster),
		open:         make(chan *Listener, 100),
		close:        make(chan *Listener, 100),
		delete:       make(chan string, 100),
		messages:     make(chan *entities.Message, 100),
	}

	go RoomManager.run()
	return RoomManager
}

func (m *Manager) run() {
	for {
		select {
		case listener := <-m.open:
			m.register(listener)
		case listener := <-m.close:
			m.deregister(listener)
		case roomid := <-m.delete:
			m.deleteBroadcast(roomid)
		case message := <-m.messages:
			fmt.Println("send to m.messages")
			m.room(message.RoomId).Submit(message.UserId + ": " + message.Text)
		}
	}
}

func (m *Manager) register(listener *Listener) {
	m.room(listener.RoomId).Register(listener.Chan)
}

func (m *Manager) deregister(listener *Listener) {
	m.room(listener.RoomId).Unregister(listener.Chan)
	close(listener.Chan)
}

func (m *Manager) deleteBroadcast(roomid string) {
	b, ok := m.roomChannels[roomid]
	if ok {
		b.Close()
		delete(m.roomChannels, roomid)
	}
}

func (m *Manager) room(roomid string) broadcast.Broadcaster {
	b, ok := m.roomChannels[roomid]
	if !ok {
		b = broadcast.NewBroadcaster(10)
		m.roomChannels[roomid] = b
	}
	return b
}

func (m *Manager) OpenListener(roomid string) chan interface{} {
	listener := make(chan interface{})
	m.open <- &Listener{
		RoomId: roomid,
		Chan:   listener,
	}
	return listener
}

func (m *Manager) CloseListener(roomid string, channel chan interface{}) {
	m.close <- &Listener{
		RoomId: roomid,
		Chan:   channel,
	}
}

func (m *Manager) DeleteBroadcast(roomid string) {
	m.delete <- roomid
}

func (m *Manager) Submit(userid, roomid, text string) {
	msg := &entities.Message{
		UserId: userid,
		RoomId: roomid,
		Text:   text,
	}
	m.messages <- msg
}
