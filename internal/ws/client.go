package ws

import (
	"encoding/json"
	"io"
	"net"
	"sync"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/jkuri/urtmp/pkg/lib"
)

// client defines websocket connection. It contains
// logic of receiving and sending the messages.
type client struct {
	mu   sync.Mutex
	conn io.ReadWriteCloser
	c    net.Conn
	subs []string
}

// receive reads next message from user's underlying connection,
// it blocks until full message is received.
func (c *client) receive() (*message, error) {
	msg, err := c.readMessage()
	if err != nil {
		c.conn.Close()
		return nil, err
	}
	if msg == nil {
		// handled some controle message.
		return nil, nil
	}

	return msg, nil
}

// send sends message to user's underlying connection.
func (c *client) send(mtype string, data map[string]interface{}) error {
	return c.write(message{Type: mtype, Data: data})
}

func (c *client) subscribe(sub string) {
	if !c.isSubscribed(sub) {
		c.subs = append(c.subs, sub)
	}
}

func (c *client) unsubscribe(sub string) {
	for i, s := range c.subs {
		if s == sub {
			c.subs = append(c.subs[:i], c.subs[i+1:]...)
			break
		}
	}
}

func (c *client) isSubscribed(data string) bool {
	for _, s := range c.subs {
		if s == data {
			return true
		}
	}
	return false
}

func (c *client) readMessage() (*message, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	h, r, err := wsutil.NextReader(c.conn, ws.StateServerSide)
	if err != nil {
		return nil, err
	}
	if h.OpCode.IsControl() {
		return nil, wsutil.ControlFrameHandler(c.conn, ws.StateServerSide)(h, r)
	}

	msg := &message{}
	if err := lib.DecodeJSON(r, msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (c *client) write(data interface{}) error {
	w := wsutil.NewWriter(c.conn, ws.StateServerSide, ws.OpText)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}

	return w.Flush()
}
