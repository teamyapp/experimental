package channel

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type OnMessageArriveCallback func(message string)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Channel interface {
	SendMessage(message string)
	OnMessageArrive(callback OnMessageArriveCallback)
	Disconnect()
}

type WebSocketChannel struct{
	conn	*websocket.Conn
}

func (w WebSocketChannel) Disconnect() {
	w.conn.Close()
}

func (w WebSocketChannel) SendMessage(message string) {
	w.conn.NextWriter(websocket.TextMessage)

	writer, err := w.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		log.Fatal(err)
		return
	}

	writer.Write([]byte(message))

	if err = writer.Close(); err != nil {
		log.Fatal(err)
		return
	}
}

func (w WebSocketChannel) OnMessageArrive(callback OnMessageArriveCallback) {
	panic("implement me")
}

func NewWebSocketChannel(w http.ResponseWriter, r *http.Request) *WebSocketChannel {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return nil
	}

	socket := &WebSocketChannel{
		conn: conn,
	}

	return socket
}


var _ Channel = (*WebSocketChannel)(nil)

