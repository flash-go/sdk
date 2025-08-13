package state

import (
	"log"
	"net/http"

	"github.com/flash-go/flash/state"
	"github.com/hashicorp/consul/api"
)

type InsecureAuthConfig struct {
	Address string
	Token   string
}

type SecureAuthConfig struct {
	Address            string
	CAPem              []byte
	CertPEM            []byte
	KeyPEM             []byte
	InsecureSkipVerify bool
	Token              string
}

// Create insecure state service without auth token
func NewWithoutAuth(consulAddress string) state.State {
	// Create state config
	config := api.DefaultConfig()
	config.Address = consulAddress

	// Create state service
	return New(config)
}

// Create insecure state service with auth token
func NewWithInsecureAuth(insecureAuthConfig *InsecureAuthConfig) state.State {
	// Create state config
	config := api.DefaultConfig()
	config.Address = insecureAuthConfig.Address
	config.Token = insecureAuthConfig.Token

	// Create state service
	return New(config)
}

// Create secure state service with auth token
func NewWithSecureAuth(secureAuthConfig *SecureAuthConfig) state.State {
	// Create state config
	config := api.DefaultConfig()
	config.Address = secureAuthConfig.Address
	config.Token = secureAuthConfig.Token
	config.Scheme = "https"

	// Create consul tls config
	consulTLSConfig, err := api.SetupTLSConfig(
		&api.TLSConfig{
			Address:            secureAuthConfig.Address,
			CAPem:              secureAuthConfig.CAPem,
			CertPEM:            secureAuthConfig.CertPEM,
			KeyPEM:             secureAuthConfig.KeyPEM,
			InsecureSkipVerify: secureAuthConfig.InsecureSkipVerify,
		},
	)
	if err != nil {
		log.Fatalf("failed to create consul tls config: %v", err)
	}

	// Create http client
	config.HttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: consulTLSConfig,
		},
	}

	// Create state service
	return New(config)
}

// Create state service
func New(config *api.Config) state.State {
	stateService, err := state.New(config)
	if err != nil {
		log.Fatalf("failed to create state service: %v", err)
	}
	return stateService
}
