package codec

import "context"

type Msg interface {
	WithClientRPCName(string)
	WithCalleeServiceName(string)
	WithCalleeApp(string)
	WithCalleeServer(string)
	WithCalleeService(string)
	WithSerializationType(string)
	WithCalleeMethod(string)
	GetAction() string
}

type msg struct {
	action string
}

const (
	ContextKeyMessage   = "message"
	ContextKeyRequestId = "RequestId"
	SerializationTypePB = "pb"
)

func (msg) WithClientRPCName(rpc string) {}

func (msg) WithCalleeServiceName(string) {}
func (msg) WithCalleeApp(string)         {}
func (msg) WithCalleeServer(string)      {}
func (msg) WithCalleeService(string)     {}
func (m *msg) WithCalleeMethod(method string) {
	m.action = method
}
func (msg) WithSerializationType(string) {}
func (m msg) GetAction() string {
	return m.action
}

func WithCloneMessage(ctx context.Context) (context.Context, Msg) {
	newMsg := &msg{}
	ctx = context.WithValue(ctx, ContextKeyMessage, newMsg)
	return ctx, newMsg
}

func PutBackMessage(Msg) {}
