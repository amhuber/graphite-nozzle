package metrics

import "time"

type StatsdClient interface {
	Gauge(stat string, value int64) error
	FGauge(stat string, value float64) error
	Incr(stat string, count int64) error
	Timing(string, int64) error
	PrecisionTiming(stat string, delta time.Duration) error
}

type Metric interface {
	Send(StatsdClient) error
}

type CounterMetric struct {
	Stat  string
	Value int64
}

type GaugeMetric struct {
	Stat  string
	Value int64
}

type FGaugeMetric struct {
	Stat  string
	Value float64
}

type TimingMetric struct {
	Stat  string
	Value int64
}

type PrecisionTimingMetric struct {
	Stat  string
	Value time.Duration
}

func (m CounterMetric) Send(statsdClient StatsdClient) error {
	statsdClient.Incr(m.Stat, m.Value)
	return nil
}

func (m GaugeMetric) Send(statsdClient StatsdClient) error {
	statsdClient.Gauge(m.Stat, m.Value)
	return nil
}

func (m FGaugeMetric) Send(statsdClient StatsdClient) error {
	statsdClient.FGauge(m.Stat, m.Value)
	return nil
}

func (m TimingMetric) Send(statsdClient StatsdClient) error {
	statsdClient.Timing(m.Stat, m.Value)
	return nil
}

func (m PrecisionTimingMetric) Send(statsdClient StatsdClient) error {
	statsdClient.PrecisionTiming(m.Stat, m.Value)
	return nil
}
