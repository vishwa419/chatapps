package main

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	id   string
	sock *websocket.Conn
	send chan []byte
	hub  *Hub
}

type Message struct {
	ClientID string `json:"clientID"`
	Text     string `json:"text"`
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// upgrade the websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	id := uuid.New()
	// make a client and register them in the hub
	c := &Client{id: id.String(), sock: conn, hub: hub, send: make(chan []byte)}
	hub.register <- c
	//read from hub and send to client
	go read(c)
	//write to broadcast channel
	go texter(c)
}

func read(c *Client) {
	defer func() {
		c.hub.unregister <- c
		c.sock.Close()
	}()

	c.sock.SetReadLimit(maxMessageSize)
	c.sock.SetReadDeadline(time.Now().Add(pongWait))
	c.sock.SetPongHandler(func(string) error {
		c.sock.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msg, err := c.sock.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		m := &Message{}
		reader := bytes.NewReader(msg)
		decoder := json.NewDecoder(reader)
		err = decoder.Decode(m)
		if err != nil {
			log.Println(err)
			return
		}
		//log.Println()
		c.hub.broadcast <- &Message{ClientID: c.id, Text: m.Text}

	}
}

func texter(c *Client) {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.sock.Close()
	}()
	// read the msgs which are there in send channel of the client put by hub, which reads data from socket
	for {
		select {
		case msg, ok := <-c.send:
			c.sock.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.sock.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.sock.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(msg)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(msg)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.sock.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.sock.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}

		}
	}

}
