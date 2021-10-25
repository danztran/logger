package logger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLogger_Infod(t *testing.T) {
	log := MustNamed("test")
	defer log.Infod()("Test log infod should take about 3s")
	time.Sleep(3 * time.Second)
}

func TestWarnd(t *testing.T) {
	log := MustNew()

	t.Run("log duration", func(t *testing.T) {
		defer log.Warnd(10 * time.Second)("this should not be displayed at all")
		time.Sleep(300 * time.Millisecond)
	})

	t.Run("log warn", func(t *testing.T) {
		defer log.Warnd(10 * time.Millisecond)("this should be displayed in warn level")
		time.Sleep(30 * time.Millisecond)
	})

	t.Run("log auto", func(t *testing.T) {
		defer log.Autod(50 * time.Millisecond)("this should be logged in warn")
		defer log.Autod(300 * time.Millisecond)("this should be logged in debug")
		time.Sleep(100 * time.Millisecond)
	})
}

func Benchmark_Zap(b *testing.B) {
	log, err := NewZap("benzap")
	require.NoError(b, err)
	for i := 0; i < b.N; i++ {
		log.Debug("benchmarking")
	}
}

func Benchmark_Logger(b *testing.B) {
	log := MustNamed("benlogger")
	for i := 0; i < b.N; i++ {
		log.Debug("benchmarking")
	}
}
