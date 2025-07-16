package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"go-plugin-example/shared" // Import the shared package
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "host",
		Output: os.Stderr,
		Level:  hclog.Debug,
	})

	// We're a host! Start by launching the plugin here.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: shared.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"greeter": &shared.GreeterPlugin{}, // The host provides a *template* for the plugin
		},
		Cmd:    exec.Command("../plugin/greeter_plugin.plug"), // Path to the compiled plugin binary
		Logger: logger,
	})
	defer client.Kill() // Ensure the plugin process is killed when the host exits

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatalf("Error connecting to plugin: %s", err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("greeter")
	if err != nil {
		log.Fatalf("Error dispensing plugin: %s", err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	greeter, ok := raw.(shared.Greeter)
	if !ok {
		log.Fatalf("Expected plugin to be a Greeter")
	}

	fmt.Println(greeter.Greet()) // Call the plugin's method
}
