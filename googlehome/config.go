package googlehome

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	defaultPort   = 8009
	defaultLang   = "en"
	defaultAccent = "us"
)

// Config has setting parameters for Client.
type Config struct {
	Context  context.Context
	Hostname string
	Port     int
	Lang     string
	Accent   string
}

// GetOrCreateContext gets context if context is set.
// If context is nil, then creates new context.
func (c Config) GetOrCreateContext() context.Context {
	if c.Context != nil {
		return c.Context
	}
	return context.Background()
}

// GetIPv4 returns IPv4 address from Hostname.
func (c Config) GetIPv4() (net.IP, error) {
	host := c.GetHostname()
	// Use IPv4 addr
	if ip := net.ParseIP(host); ip != nil {
		return ip, nil
	}

	// Use hostname
	ips, err := net.LookupIP(host)
	switch {
	case err != nil:
		return nil, err
	case len(ips) == 0:
		return nil, fmt.Errorf("ip addr is empty for the hostname: %s", host)
	}
	return ips[0], nil
}

// GetHostname returns Hostname.
func (c Config) GetHostname() string {
	if c.Hostname != "" {
		return c.Hostname
	}
	return os.Getenv("GOOGLE_HOME_HOST")
}

// GetPort returns port number.
func (c Config) GetPort() int {
	if c.Port > 0 {
		return c.Port
	}
	port, _ := strconv.Atoi(os.Getenv("GOOGLE_HOME_PORT"))
	if port > 0 {
		return port
	}
	return defaultPort
}

// GetLang returns speaking language.
func (c Config) GetLang() string {
	if c.Lang != "" {
		return c.Lang
	}
	lang := os.Getenv("GOOGLE_HOME_LANG")
	if lang != "" {
		return lang
	}
	return defaultLang
}

// GetAccent returns language accent.
func (c Config) GetAccent() string {
	if c.Accent != "" {
		return c.Accent
	}
	accent := os.Getenv("GOOGLE_HOME_ACCENT")
	if accent != "" {
		return accent
	}
	return defaultAccent
}
