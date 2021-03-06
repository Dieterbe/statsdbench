package statsdbench

import (
	"fmt"
	"net"
	"strconv"
	"testing"
	"time"

	dieterbe "github.com/Dieterbe/statsd-go"
	alexcesaro "github.com/alexcesaro/statsd"
	cactus "github.com/cactus/go-statsd-client/statsd"
	"github.com/peterbourgon/g2s"
	quipo "github.com/quipo/statsd"
	"github.com/raintank/met/helper"
)

const (
	addr        = ":8125"
	prefix      = "prefix."
	prefixNoDot = "prefix"
	counterKey  = "foo.bar.counter"
	gaugeKey    = "foo.bar.gauge"
	gaugeValue  = 42
	timingKey   = "foo.bar.timing"
	tValDur     = 153 * time.Millisecond
	tValInt     = int(153)
	tValInt64   = int64(153)
	flushPeriod = 100 * time.Millisecond
)

type logger struct{}

func (logger) Println(v ...interface{}) {}

func BenchmarkAlexcesaro(b *testing.B) {
	s := newServer()
	c, err := alexcesaro.New(addr, alexcesaro.WithPrefix(prefix), alexcesaro.WithFlushPeriod(flushPeriod))
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Increment(counterKey)
		c.Gauge(gaugeKey, gaugeValue)
		c.Timing(timingKey, tValInt, 1)
	}
	c.Close()
	b.StopTimer()
	s.Close()
}

func BenchmarkCactus(b *testing.B) {
	s := newServer()
	c, err := cactus.NewBufferedClient(addr, prefix, flushPeriod, 1432)
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Inc(counterKey, 1, 1)
		c.Gauge(gaugeKey, gaugeValue, 1)
		c.Timing(timingKey, tValInt64, 1)
	}
	c.Close()
	b.StopTimer()
	s.Close()
}

func BenchmarkCactusTimingAsDuration(b *testing.B) {
	s := newServer()
	c, err := cactus.NewBufferedClient(addr, prefix, flushPeriod, 1432)
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Inc(counterKey, 1, 1)
		c.Gauge(gaugeKey, gaugeValue, 1)
		c.TimingDuration(timingKey, tValDur, 1)
	}
	c.Close()
	b.StopTimer()
	s.Close()
}

func BenchmarkDieterbe(b *testing.B) {
	s := newServer()
	c, err := dieterbe.NewClient(true, addr, prefix)
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Increment(counterKey)
		c.Gauge(gaugeKey, gaugeValue)
		c.Timing(timingKey, tValInt64)
	}
	c.Close()
	b.StopTimer()
	s.Close()
}

func BenchmarkDieterbeRaw(b *testing.B) {
	s := newServer()
	c, err := dieterbe.NewClient(true, addr, prefix)
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()
	counter := []byte(fmt.Sprintf("%s:1|c", counterKey))
	gauge := []byte(fmt.Sprintf("%s:%d|g", gaugeKey, gaugeValue))
	timing := []byte(fmt.Sprintf("%s:%d|ms", timingKey, tValInt))
	for i := 0; i < b.N; i++ {
		c.SendRaw(counter)
		c.SendRaw(gauge)
		c.SendRaw(timing)
	}
	c.Close()
	b.StopTimer()
	s.Close()
}

func BenchmarkG2s(b *testing.B) {
	s := newServer()
	c, err := g2s.Dial("udp", addr)
	if err != nil {
		b.Fatal(err)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Counter(1, counterKey, 1)
		c.Gauge(1, gaugeKey, strconv.Itoa(gaugeValue))
		c.Timing(1, timingKey, tValDur)
	}
	b.StopTimer()
	s.Close()
}

func BenchmarkHelperCesaro(b *testing.B) {
	s := newServer()
	c, err := helper.New(true, addr, "standard", prefix, "")
	if err != nil {
		b.Fatal(err)
	}
	counter := c.NewCount(counterKey)
	gauge := c.NewGauge(gaugeKey, 0)
	timing := c.NewTimer(timingKey, 0)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		counter.Inc(1)
		gauge.Value(gaugeValue)
		timing.Value(tValDur)
	}
	b.StopTimer()
	s.Close()
}

func BenchmarkQuipo(b *testing.B) {
	s := newServer()
	c := quipo.NewStatsdBuffer(flushPeriod, quipo.NewStatsdClient(addr, prefix))
	c.Logger = logger{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Incr(counterKey, 1)
		c.Gauge(gaugeKey, gaugeValue)
		c.Timing(timingKey, tValInt64)
	}
	c.Close()
	b.StopTimer()
	s.Close()
}

func BenchmarkQuipoTimingAsDuration(b *testing.B) {
	s := newServer()
	c := quipo.NewStatsdBuffer(flushPeriod, quipo.NewStatsdClient(addr, prefix))
	c.Logger = logger{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Incr(counterKey, 1)
		c.Gauge(gaugeKey, gaugeValue)
		c.PrecisionTiming(timingKey, tValDur)
	}
	c.Close()
	b.StopTimer()
	s.Close()
}

type server struct {
	conn   *net.UDPConn
	closed chan bool
}

func newServer() *server {
	addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	s := &server{conn: conn, closed: make(chan bool)}
	go func() {
		buf := make([]byte, 512)
		for {
			_, err := conn.Read(buf)
			if err != nil {
				s.closed <- true
				return
			}
		}
	}()
	return s
}

func (s *server) Close() {
	s.conn.Close()
	<-s.closed
	return
}
