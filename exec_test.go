package main

import (
	"os"
	"testing"

	"github.com/bmizerany/assert"
)

func TestPrepare(t *testing.T) {
	//set env
	os.Setenv(ENV_SWARM_ENDPOINTS, "52.78.18.255:3376,52.78.19.214:1234,52.78.35.207:5678")

	expect := []string{"52.78.18.255:3376", "52.78.19.214:1234", "52.78.35.207:5678"}
	endpoints, _, _, _ := prepare()
	assert.Equal(t, expect, endpoints)

	//unset env
	os.Unsetenv(ENV_SWARM_ENDPOINTS)
}
