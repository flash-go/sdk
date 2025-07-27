package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/flash-go/flash/state"
)

type Config interface {
	GetService() string
	Get(key string) string
	GetInt(key string) int
	Set(key, value string) error
	SetEnvMap(envMap map[string]string)
}

type config struct {
	service string
	state   state.State
}

func New(state state.State, service string) Config {
	if service == "" {
		log.Fatal("service is not set")
	}
	return &config{
		service: service,
		state:   state,
	}
}

func (c *config) GetService() string {
	return c.service
}

func (c *config) Get(key string) string {
	k := c.service + key
	v, err := c.state.GetValue(k)
	if err != nil {
		log.Fatalf("failed to get config key [%s]: %v", k, err)
	}
	return v
}

func (c *config) GetInt(key string) int {
	v, err := strconv.Atoi(c.Get(key))
	if err != nil {
		log.Fatalf("failed to parse int config key [%s]: %v", key, err)
	}
	return v
}

func (c *config) Set(key, value string) error {
	k := c.service + key
	_, err := c.state.GetValue(k)
	if errors.Is(err, state.ErrKeyNotFound) {
		return c.state.SetValue(k, value)
	}
	return err
}

func (c *config) SetEnvMap(envMap map[string]string) {
	for env, key := range envMap {
		if err := c.Set(key, os.Getenv(env)); err != nil {
			log.Fatalf("failed to create KV [%s]: %v", key, err)
		}
	}
}
