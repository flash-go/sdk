package state

import (
	"log"

	"github.com/flash-go/flash/state"
	"github.com/hashicorp/consul/api"
)

// Create state service
func New(consulAddress string) state.State {
	// Create state config
	config := api.DefaultConfig()
	config.Address = consulAddress

	// Create state service
	stateService, err := state.New(config)
	if err != nil {
		log.Fatalf("failed to create state service: %v", err)
	}

	return stateService
}
