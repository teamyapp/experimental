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
	SendMessage(message string) error
	OnMessageReceived() chan []byte
	Disconnect()
	Listen()
}

type WebSocketChannel struct{
	conn		*websocket.Conn
	onMessageReceived chan []byte
}

func (w WebSocketChannel) Listen() {
	go func() {
		for {
			_, data, err := w.conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("error: %v", err)
				}
				break
			}
			w.onMessageReceived <- data
		}
	}()
}

func (w WebSocketChannel) Disconnect() {
	w.conn.Close()
}

func (w WebSocketChannel) SendMessage(message string) error {
	writer, err := w.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		log.Fatal(err)
		return err
	}

	writer.Write([]byte(message))

	if err = writer.Close(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (w WebSocketChannel) OnMessageReceived() chan []byte {
	return w.onMessageReceived
}

func NewWebSocketChannel(w http.ResponseWriter, r *http.Request) (*WebSocketChannel, error) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return nil, err
	}

	socket := &WebSocketChannel{
		conn: conn,
		onMessageReceived: make(chan []byte, 0),
	}

	return socket, nil
}


var _ Channel = (*WebSocketChannel)(nil)

