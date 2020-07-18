package options

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLogging(t *testing.T) {
	assert := assert.New(t)

	Init()
	assert.Equal(logrus.WarnLevel, logrus.GetLevel(), "Default log level should be warn")

	os.Setenv("LOGGING_LEVEL", "info")
	Init()
	assert.Equal(logrus.InfoLevel, logrus.GetLevel())
}
