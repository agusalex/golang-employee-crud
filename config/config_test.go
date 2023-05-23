package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigFailFetch(t *testing.T) {
	defer func() { recover() }()
	Value = nil
	Get()
	t.Errorf("should have panicked")
}

func TestDBData(t *testing.T) {
	FetchConfig()
	var configuration = Get()
	assert.Equal(t, "0.0.0.0", configuration.DB.Host)
	assert.Equal(t, "root", configuration.DB.User)
	assert.Equal(t, "mydbname", configuration.DB.Name)
	assert.Equal(t, "my-secret-pw", configuration.DB.Password)
}

func TestHost(t *testing.T) {
	FetchConfig()
	var configuration = Get()
	assert.Equal(t, "8080", configuration.Server.Port)
}
