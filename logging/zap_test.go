package logging

import (
	"testing"
)

func TestZapLogger(t *testing.T) {
	Init("../log", "test.log", true)
	logger.Info("test")
}
