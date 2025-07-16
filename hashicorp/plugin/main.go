package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"go-plugin-example/shared" // Import the shared package
)

// GreeterImpl is the concrete implementation of the Greeter interface within the plugin.
type GreeterImpl struct{}

func (g *GreeterImpl) Greet() string {
	return "Hello from the Go-Plugin!"
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "greeter-plugin",
		Output: os.Stderr, // Plugin logs usually go to Stderr, which is mirrored by the host
		Level:  hclog.Debug,
	})

	pluginMap := map[string]plugin.Plugin{
		"greeter": &shared.GreeterPlugin{Impl: &GreeterImpl{}},
	}

	logger.Debug("Starting plugin server...")

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins:         pluginMap,
		Logger:          logger,
	})
}
