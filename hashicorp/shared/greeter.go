package shared

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// Greeter is the interface that we're exposing as a plugin.
type Greeter interface {
	Greet() string
}

// Here is an implementation that writes to stdout and uses the RPC client to get it back.
// This is the implementation that the plugin will serve.
type GreeterRPC struct{ client *rpc.Client }

func (g *GreeterRPC) Greet() string {
	var resp string
	err := g.client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		// You would typically handle errors more robustly
		return "Error calling plugin: " + err.Error()
	}
	return resp
}

// Here is the RPC server that GreeterRPCClient talks to.
type GreeterRPCServer struct {
	Impl Greeter
}

func (s *GreeterRPCServer) Greet(args interface{}, resp *string) error {
	*resp = s.Impl.Greet()
	return nil
}

// GreeterPlugin is the plugin.Plugin implementation for our Greeter.
type GreeterPlugin struct {
	Impl Greeter
}

func (p *GreeterPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &GreeterRPCServer{Impl: p.Impl}, nil
}

func (p *GreeterPlugin) Client(broker *plugin.MuxBroker, client *rpc.Client) (interface{}, error) {
	return &GreeterRPC{client: client}, nil
}

// HandshakeConfig is used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion: 1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}
