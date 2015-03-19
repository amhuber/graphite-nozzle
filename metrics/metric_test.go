package metrics_test

import (
	. "github.com/CloudCredo/graphite-nozzle/metrics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"
)

type FakeStatsdClient struct {
	timingCalled          bool
	precisionTimingCalled bool
	incrCalled            bool
	gaugeCalled           bool
	fGaugeCalled          bool
	stat                  string
	value                 int64
	fValue                float64
	precisionTimingValue  time.Duration
}

func (f *FakeStatsdClient) Timing(stat string, delta int64) error {
	f.timingCalled = true
	f.stat = stat
	f.value = delta
	return nil
}

func (f *FakeStatsdClient) PrecisionTiming(stat string, delta time.Duration) error {
	f.precisionTimingCalled = true
	f.stat = stat
	f.precisionTimingValue = delta
	return nil
}

func (f *FakeStatsdClient) Incr(stat string, count int64) error {
	f.incrCalled = true
	f.stat = stat
	f.value = count
	return nil
}

func (f *FakeStatsdClient) Gauge(stat string, value int64) error {
	f.gaugeCalled = true
	f.stat = stat
	f.value = value
	return nil
}

func (f *FakeStatsdClient) FGauge(stat string, value float64) error {
	f.fGaugeCalled = true
	f.stat = stat
	f.fValue = value

	return nil
}

var _ = Describe("Metric", func() {
	var (
		fakeStatsdClient *FakeStatsdClient
	)

	Describe("#Send", func() {
		Context("with a TimingMetric", func() {
			It("sends the Metric to StatsD with int64 precision", func() {
				fakeStatsdClient = new(FakeStatsdClient)
				metric := TimingMetric{
					Stat:  "http.responsetimes.api_10_244_0_34_xip_io",
					Value: 10,
				}

				metric.Send(fakeStatsdClient)
				Expect(fakeStatsdClient.timingCalled).To(BeTrue())
				Expect(fakeStatsdClient.stat).To(Equal("http.responsetimes.api_10_244_0_34_xip_io"))
				Expect(fakeStatsdClient.value).To(Equal(int64(10)))
			})
		})

		Context("with a PrecisionTimingMetric", func() {
			It("sends the Metric to StatsD with time.Duration precision", func() {
				fakeStatsdClient = new(FakeStatsdClient)
				metric := PrecisionTimingMetric{
					Stat:  "http.responsetimes.api_10_244_0_34_xip_io",
					Value: 50 * time.Millisecond,
				}

				metric.Send(fakeStatsdClient)
				Expect(fakeStatsdClient.precisionTimingCalled).To(BeTrue())
				Expect(fakeStatsdClient.stat).To(Equal("http.responsetimes.api_10_244_0_34_xip_io"))
				Expect(fakeStatsdClient.precisionTimingValue).To(Equal(50 * time.Millisecond))
			})
		})

		Context("with a CounterMetric", func() {
			It("sends the Metric to StatsD with int64 precision", func() {
				fakeStatsdClient = new(FakeStatsdClient)
				metric := CounterMetric{
					Stat:  "http.statuscodes.api_10_244_0_34_xip_io.200",
					Value: 1,
				}

				metric.Send(fakeStatsdClient)
				Expect(fakeStatsdClient.incrCalled).To(BeTrue())
				Expect(fakeStatsdClient.stat).To(Equal("http.statuscodes.api_10_244_0_34_xip_io.200"))
				Expect(fakeStatsdClient.value).To(Equal(int64(1)))
			})
		})

		Context("with a GaugeMetric", func() {
			It("sends the Metric to StatsD with int64 precision", func() {
				fakeStatsdClient = new(FakeStatsdClient)
				metric := GaugeMetric{
					Stat:  "router__0.numCPUS",
					Value: 4,
				}

				metric.Send(fakeStatsdClient)
				Expect(fakeStatsdClient.gaugeCalled).To(BeTrue())
				Expect(fakeStatsdClient.stat).To(Equal("router__0.numCPUS"))
				Expect(fakeStatsdClient.value).To(Equal(int64(4)))
			})
		})

		Context("with an FGaugeMetric", func() {
			It("sends the Metric to StatsD with float64 precision", func() {
				fakeStatsdClient = new(FakeStatsdClient)
				metric := FGaugeMetric{
					Stat:  "router__0.numCPUS",
					Value: 4,
				}

				metric.Send(fakeStatsdClient)
				Expect(fakeStatsdClient.fGaugeCalled).To(BeTrue())
				Expect(fakeStatsdClient.stat).To(Equal("router__0.numCPUS"))
				Expect(fakeStatsdClient.fValue).To(Equal(float64(4)))
			})
		})
	})
})
