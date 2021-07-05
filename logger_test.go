package logger

import (
	"testing"
	"time"
)

func TestLogger_Infod(t *testing.T) {
	log := MustNew("test logger")
	defer log.Infod("Test log infod should take about 3s")()
	time.Sleep(3 * time.Second)
}
