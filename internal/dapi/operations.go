package dapi

// Opcode distinguishes ints as Opcodes.
type Opcode uint8

const (
	// OpIdentify is for droplet authentication.
	OpIdentify Opcode = iota
	// OpHello is for signaling a droplet to identify
	OpHello
	// OpRegisterListener is for registering new listeners.
	OpRegisterListener
	// OpRegisterCommand is for registering new commands.
	OpRegisterCommand
)

// IdentifyMessage holds information to be sent by droplet for identifying.
type IdentifyMessage struct {
}

// NewIdentifyMessage creates a new NetMessage with appropriate data for a
// Identify message.
func NewIdentifyMessage() *NetMessage {
	// Currently we have nothing for the identify message.
	msg := &NetMessage{
		Op:   OpIdentify,
		Data: &IdentifyMessage{},
	}
	return msg
}

// HelloMessage holds information to be sent to a droplet for initiating a
// connection.
type HelloMessage struct {
}

// NewHelloMessage creates a new NetMessage with appropriate data for a Hello
// message.
func NewHelloMessage() *NetMessage {
	msg := &NetMessage{
		Op:   OpHello,
		Data: &HelloMessage{},
	}
	return msg
}

// RegisterListenerMessage holds information sent by a droplet to register a
// listener.
type RegisterListenerMessage struct {
}

// RegisterCommandMessage holds information sent by a droplet to register a
// listener.
type RegisterCommandMessage struct {
}
