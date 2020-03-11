package cfg

import (
	"fmt"
	"sync"
)

// Config set the configuration from different source: file, environment variable, docker secrets
type Config struct {
	m  map[string]string
	mu sync.RWMutex
}

// Provider is implemente by the user to provide configuration as one map
type Provider interface {
	Provide() (map[string]string, error)
}

// New constructor function got new Config from provider. it will return an error if there is a problem
func New(p Provider) (*Config, error) {
	m, err := p.Provide()
	if err != nil {
		return nil, err
	}
	c := &Config{
		m: m,
	}
	return c, nil
}

// String return the value of a given key as string, it will return an error if not found
func (c *Config) String(key string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, found := c.m[key]
	if !found {
		return "", fmt.Errorf("Unknow key: %v ", key)
	}
	return value, nil
}

// MustString return the value of the given key as string, it will panic is there is an error
// if will get the map[string]string and get the value for a given key
func (c *Config) MustString(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, found := c.m[key]
	if !found {
		panic(fmt.Sprintf("Unknow key %s : ", key))
	}
	return value

}
