package dapi

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/gorilla/websocket"
)

// NetErrorType is used to define errors for a NetHelper
type NetErrorType int

const (
	// BadRead means we had an error reading from websocket.Conn.
	// BadRead and BadWrite are accounted as unknown errors.
	BadRead NetErrorType = (4000 + iota)
	// BadWrite means we had an error writing to websocket.Conn
	// BadRead and BadWrite are accounted as unknown errors.
	BadWrite
	// BadPack means we couldn't pack a NetMessage for writing.
	BadPack
	// BadUnpack means we couldn't unpack a read message from websocket.Conn.
	// Usually means the droplet sent a bad message.
	BadUnpack
	// BadMessage means a droplet sent an incorrect message type.
	BadMessage
	// BadIdentify means a droplet sent a message before identifying.
	BadIdentify
	// IdentifyTimeout means a droplet failed to identify in time.
	IdentifyTimeout
	// ConnectionClosed means a droplet closed the connection.
	ConnectionClosed
	// Timeout is used to indicated a write or read exceeded its deadline.
	Timeout
)

// NetError is used to create Errors specific to this domain.
type NetError struct {
	ErrType NetErrorType
	Message string
}

// Error creates an error string
func (e *NetError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

// NetMessage defines the type of JSON to send over a websocket to a droplet.
// It is also used to determine the types of incoming messages.
type NetMessage struct {
	Op   Opcode
	Data interface{}
}

// The NetHelper struct provides functions for handling raw binary data that comes
// any websockets tied to the DAPI.
type NetHelper struct {
	conn *websocket.Conn
}

// NewNetHelper creates a new Net struct for handling Net specific messages.
func NewNetHelper(c *websocket.Conn) *NetHelper {
	net := &NetHelper{
		conn: c,
	}
	return net
}

// Pack takes a struct and encodes it into JSON.
func (n *NetHelper) Pack(netmsg *NetMessage) ([]byte, error) {
	jsonData, err := json.Marshal(netmsg)
	if err != nil {
		return nil, &NetError{BadPack,
			fmt.Sprintf("unable to pack net message: %s", err.Error())}
	}
	return jsonData, nil
}

// Unpack takes a JSON message and decodes it into the given struct.
func (n *NetHelper) Unpack(data []byte) (*NetMessage, error) {
	netmsg := &NetMessage{}
	err := json.Unmarshal(data, &netmsg)
	if err != nil {
		return nil, &NetError{BadUnpack,
			fmt.Sprintf("unable to unpack net message: %s", err.Error())}
	}
	return netmsg, nil
}

// Conn returns the intenral connection of the NetHelper. Only use this when
// the other functions provided by NetHelper do not satisfy your needs.
func (n *NetHelper) Conn() *websocket.Conn {
	return n.conn
}

// Write will write a NetMessage struct to it internal websocket connection.
func (n *NetHelper) Write(netmsg *NetMessage) error {
	jsonmsg, err := n.Pack(netmsg)
	if err != nil {
		return &NetError{BadPack,
			fmt.Sprintf("unable to pack net message: %s", err.Error())}
	}
	err = n.conn.WriteMessage(websocket.BinaryMessage, jsonmsg)
	if err != nil {
		return &NetError{BadWrite,
			fmt.Sprintf("unable to write net message: %s", err.Error())}
	}
	return nil
}

// Read will read an incoming message from the internal websocket connection
// and will attempt to return an appropriate net message.
func (n *NetHelper) Read(t time.Duration) (*NetMessage, error) {
	if t != 0 {
		n.conn.SetReadDeadline(time.Now().Add(t))
	}
	mt, msg, err := n.conn.ReadMessage()
	if err != nil {
		netErr, ok := err.(net.Error)
		if ok && netErr.Timeout() {
			return nil, &NetError{Timeout,
				fmt.Sprintf("read deadline exceeded : %s", err.Error())}
		}
		return nil, &NetError{BadRead,
			fmt.Sprintf("failed to read net message: %s", err.Error())}
	}
	if mt != websocket.BinaryMessage {
		return nil, &NetError{BadMessage, "bad message sent"}
	}
	netmsg, err := n.Unpack(msg)
	if err != nil {
		return nil, &NetError{BadUnpack,
			fmt.Sprintf("unable to unpack net message: %s", err.Error())}
	}
	return netmsg, nil
}

// WriteCloseMessage will write a close normal to the websocket connection.
func (n *NetHelper) WriteCloseMessage() error {
	err := n.conn.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(
			websocket.CloseNormalClosure, "connection closing"))
	if err != nil {
		return err
	}
	return nil
}

// WriteCloseErrMessage will write a closing message to the droplet.
func (n *NetHelper) WriteCloseErrMessage(ne *NetError) error {
	err := n.conn.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(int(ne.ErrType), ne.Error()))
	if err != nil {
		return err
	}
	return nil
}
