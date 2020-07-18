package options

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestREST(t *testing.T) {
	assert := assert.New(t)

	Init("rest")
	assert.Equal("0.0.0.0", REST.Address, "Service addresses should always default to all interfaces")

	os.Setenv("REST_ADDRESS", "1.2.3.4")
	os.Setenv("REST_PORT", "1234")
	Init("rest")
	assert.Equal("1.2.3.4", REST.Address)
	assert.Equal(1234, REST.Port)
}
